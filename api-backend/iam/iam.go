package iam

import (
	"github.com/johncave/podinate/api-backend/account"
	"github.com/johncave/podinate/api-backend/apierror"
	"github.com/johncave/podinate/api-backend/config"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/gobwas/glob"
	eh "github.com/johncave/podinate/api-backend/errorhandler"
)

func RequestorCan(requestor Resource, account account.Account, resource Resource, action string) bool {
	// Check if the user has permission to do the action on the resource
	policies, err := GetPolicies(account, requestor)
	if err != nil {
		eh.Log.Errorw("Error getting policies for user", "error", err, "requestor", requestor.GetRID(), "account", account.GetUUID(), "account_id", account.ID)
		return false
	}

	for _, policy := range policies {
		if policy.Allows(resource, action) {
			return true
		}
	}

	eh.Log.Errorw("denying requestor this action", "requestor", requestor.GetRID(), "account", account.GetUUID(), "account_id", account.ID, "resource", resource.GetRID(), "action", action)
	return false
}

type PolicyDocument struct {
	Version    string            `yaml:"version"`
	Statements []PolicyStatement `yaml:"statements"`
}

type PolicyStatement struct {
	Effect    string   `yaml:"effect"`
	Actions   []string `yaml:"actions"`
	Resources []string `yaml:"resources"`
}

func (p *Policy) Allows(resource Resource, action string) bool {
	// Check if the policy allows the action
	var doc PolicyDocument
	err := yaml.Unmarshal([]byte(p.Content), &doc)
	if err != nil {
		eh.Log.Errorw("Error unmarshalling policy", "error", err, "policy", p.UUID)
		return false
	}

	//eh.Log.Debugw("Checking policy", "policy_uuid", p.UUID, "resource", resource.GetRID(), "action", action, "policy", p, "document", doc)
	// Check all statements for a match
	for _, statement := range doc.Statements {
		// Check if the statement allows the action
		if statement.Effect == "allow" {
			// Check if the resource matches the statement
			for _, resourcePattern := range statement.Resources {
				g := glob.MustCompile(resourcePattern, '/', ':')
				if g.Match(resource.GetRID()) {
					// Check if the action matches the statement
					for _, actionPattern := range statement.Actions {
						g = glob.MustCompile(actionPattern, '/', ':')
						if g.Match(action) {
							eh.Log.Infow("allowing action due to policy", "policy_uuid", p.UUID, "resource", resource.GetRID(), "resource_pattern", resourcePattern, "action", action, "action_pattern", actionPattern, "policy", p, "document", doc, "statement", statement)
							return true
						}
					}
				}
			}
		}
	}

	return false

}

type Resource interface {
	GetRID() string
}

func GetPolicies(account account.Account, requestor Resource) ([]Policy, error) {
	// Retrieve any policies from the policy_attachment table
	// that apply to this user and account

	rows, err := config.DB.Query("SELECT policy.uuid, policy.id, policy.content, policy.current_version FROM policy_attachment, policy WHERE policy_attachment.policy_uuid = policy.uuid AND policy_attachment.account_uuid = $1 AND policy_attachment.resource_id = $2", account.GetUUID(), requestor.GetRID())
	if err != nil {
		eh.Log.Errorw("Error retrieving policies for user", "error", err, "user", requestor.GetRID(), "account", account.GetUUID(), "account_id", account.ID)
		return []Policy{}, err
	}

	defer rows.Close()
	policies := make([]Policy, 0)
	for rows.Next() {
		var policy Policy
		err = rows.Scan(&policy.UUID, &policy.Name, &policy.Content, &policy.Version)
		if err != nil {
			eh.Log.Errorw("Error scanning policy", "error", err)
			return []Policy{}, err
		}
		policy.Account = account
		policy.Requestor = requestor
		policies = append(policies, policy)
	}

	return policies, nil
}

func CreatePolicyForAccount(account account.Account, name string, content string, versionComment string) (Policy, *apierror.ApiError) {
	// Create a new policy for the account
	var policy Policy
	policy.Account = account
	policy.Name = name
	policy.ValidateNewContent(content)

	err := config.DB.QueryRow("INSERT INTO policy(uuid, account_uuid, id, content, current_version) VALUES(gen_random_uuid(), $1, $2, $3, 1) returning uuid", account.GetUUID(), name, content).Scan(&policy.UUID)
	if err != nil {
		eh.Log.Errorw("Error creating policy", "error", err, "account", account.GetUUID(), "account_id", account.ID)
		return Policy{}, &apierror.ApiError{Code: 500, Message: "Error creating policy"}
	}

	// Insert the policy version too
	_, err = config.DB.Exec("INSERT INTO policy_version(uuid, policy_uuid, content, version_number, comment) VALUES(gen_random_uuid(), $1, $2, 1, $3)", policy.UUID, content, versionComment)
	if err != nil {
		eh.Log.Errorw("Error creating policy version", "error", err, "account", account.GetUUID(), "account_id", account.ID)
		return Policy{}, &apierror.ApiError{Code: 500, Message: "Error creating policy version"}
	}
	policy.Content = content

	return policy, nil
}

func (p *Policy) UpdatePolicy(content string, comment string) error {
	// Update the policy

	// TODO: Validate the policy content before updating it
	err := p.ValidateNewContent(content)
	if err != nil {
		eh.Log.Errorw("Error validating new policy content", "error", err, "policy", p.UUID)
		return err
	}

	err = config.DB.QueryRow("INSERT INTO policy_version(uuid, policy_uuid, content, comment, version) VALUES(gen_random_uuid(), $1, $2, $3, (SELECT policy_version + 1 FROM policy_version WHERE policy_uuid = $1 ORDER BY policy_version DESC LIMIT 1)) returning version", p.UUID, content, comment).Scan(&p.Version)
	if err != nil {
		eh.Log.Errorw("Error creating new policy version", "error", err, "policy", p.UUID)
		return err
	}

	// Update the policy version too
	_, err = config.DB.Exec("UPDATE policy SET version = $1, content = $2 WHERE uuid = $3", p.Version, content, p.UUID)
	if err != nil {
		eh.Log.Errorw("Error updating policy", "error", err, "policy", p.UUID)
		return err
	}

	p.Content = content

	return nil
}

// AttachToResource attaches a policy to a resource
func (p *Policy) AttachToResource(requestor Resource, resource Resource) *apierror.ApiError {
	// Attach the policy to the resource in the account using the policy_attachment table
	_, err := config.DB.Exec("INSERT INTO policy_attachment(policy_uuid, account_uuid, resource_id, attached_by) VALUES($1, $2, $3, $4)", p.UUID, p.Account.GetUUID(), resource.GetRID(), requestor.GetRID())
	if err != nil {
		eh.Log.Errorw("Error attaching policy to resource", "error", err, "policy", p.UUID, "resource", resource.GetRID(), "requestor", requestor.GetRID())
		return apierror.New(500, "Error attaching policy to resource")
	}
	return nil
}

func (p *Policy) ValidateNewContent(content string) error {
	// Validate that the new policy content is valid
	// TODO - implement this
	return nil
}

type Policy struct {
	UUID      string
	Account   account.Account
	Requestor Resource
	Name      string
	Content   string
	Version   int
}
