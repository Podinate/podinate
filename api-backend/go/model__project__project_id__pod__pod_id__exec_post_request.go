/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProjectProjectIdPodPodIdExecPostRequest struct {

	// The command to execute
	Command []string `json:"command"`
}

// AssertProjectProjectIdPodPodIdExecPostRequestRequired checks if the required fields are not zero-ed
func AssertProjectProjectIdPodPodIdExecPostRequestRequired(obj ProjectProjectIdPodPodIdExecPostRequest) error {
	elements := map[string]interface{}{
		"command": obj.Command,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseProjectProjectIdPodPodIdExecPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ProjectProjectIdPodPodIdExecPostRequest (e.g. [][]ProjectProjectIdPodPodIdExecPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProjectProjectIdPodPodIdExecPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProjectProjectIdPodPodIdExecPostRequest, ok := obj.(ProjectProjectIdPodPodIdExecPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProjectProjectIdPodPodIdExecPostRequestRequired(aProjectProjectIdPodPodIdExecPostRequest)
	})
}