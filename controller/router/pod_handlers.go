package router

import (
	"context"
	"log"
	"net/http"

	api "github.com/johncave/podinate/controller/go"
	"github.com/johncave/podinate/controller/iam"
	lh "github.com/johncave/podinate/controller/loghandler"
	pod "github.com/johncave/podinate/controller/pod"
	"github.com/johncave/podinate/controller/project"
	"github.com/johncave/podinate/controller/responder"
)

// PodAPIService holds all the handlers for the Pod API
type PodAPIService struct {
	api.PodApiService
}

// NewPodAPIService creates a new service for handling requests in the Pod API
func NewPodAPIService() api.PodApiServicer {
	return &PodAPIService{}
}

// ProjectProjectIdPodGet - Returns a list of pods for a project.
func (s *PodAPIService) ProjectProjectIdPodGet(ctx context.Context, projectID string, accountID string, page int32, limit int32) (api.ImplResponse, error) {

	//log.Printf("%s %s %d %d", projectID, accountID, page, limit)
	lh.Debug(ctx, "Getting pods for project", "projectID", projectID, "accountID", accountID, "page", page, "limit", limit)

	// Get the account and project
	theProject, err := getProject(ctx, accountID, projectID)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	// Get the pods for the project
	pods, apiErr := pod.GetByProject(ctx, theProject, page, limit)
	if apiErr != nil {
		return responder.Response(apiErr.Code, apiErr.Message), nil
	}

	// Convert the pods to the API format
	var apiPods []api.ProjectProjectIdPodGet200ResponseItemsInner
	for _, p := range pods {
		//lh.Debug(ctx, "Pod Environment", "env", p.Environment)
		if iam.RequestorCan(ctx, &theProject.Account, p, pod.ActionView) {
			apiPods = append(apiPods, api.ProjectProjectIdPodGet200ResponseItemsInner{
				Id:          p.ID,
				Name:        p.Name,
				ResourceId:  p.GetResourceID(),
				Image:       p.Image,
				Tag:         p.Tag,
				Status:      p.Status,
				Environment: pod.EnvVarToAPIMany(p.Environment),
				Services:    pod.ServicesToAPI(p.Services),
				Volumes:     p.Volumes.ToAPI(),
			})
		}
	}

	out := api.ProjectProjectIdPodGet200Response{
		Items: apiPods,
		Total: int32(len(apiPods)),
		Page:  page,
		Limit: limit,
	}

	if len(apiPods) == 0 {
		if !iam.RequestorCan(ctx, &theProject.Account, theProject, project.ActionView) {
			return responder.Response(http.StatusForbidden, "You do not have permission to view pods in this project"), nil
		}
	}

	return responder.Response(http.StatusOK, out), nil
}

// ProjectProjectIdPodPodIdGet - Returns a pod for a project.
func (s *PodAPIService) ProjectProjectIdPodPodIdGet(ctx context.Context, projectID string, podID string, accountID string) (api.ImplResponse, error) {

	log.Printf("%s %s %s", projectID, podID, accountID)

	theProject, err := getProject(ctx, accountID, projectID)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	p, apiErr := pod.GetByID(ctx, theProject, podID)
	if apiErr != nil {
		return responder.Response(apiErr.Code, apiErr.Message), nil
	}

	if !iam.RequestorCan(ctx, &theProject.Account, p, pod.ActionView) {
		return responder.Response(http.StatusForbidden, "You do not have permission to view this pod"), nil
	}

	lh.Debug(ctx, "Returning pod get response", "pod", p.ToAPI())

	return responder.Response(http.StatusOK, p.ToAPI()), nil
}

// ProjectProjectIdPodPodIdPut - Updates a pod for a project.
func (s *PodAPIService) ProjectProjectIdPodPodIdPut(ctx context.Context, projectID string, podID string, accountID string, podIn api.Pod) (api.ImplResponse, error) {

	lh.Debug(ctx, "Updating pod", "project_id", projectID, "acc", accountID, "podId", podID, "pod", podIn)

	// Get the account by ID
	theProject, err := getProject(ctx, accountID, projectID)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	// Get the pod by ID
	thePod, apiErr := pod.GetByID(ctx, theProject, podID)
	if apiErr != nil {
		return responder.Response(apiErr.Code, apiErr.Message), nil
	}

	lh.Debug(ctx, "Got pod", "pod", thePod, "want", podIn)

	// Check if the user can update the pod
	if !iam.RequestorCan(ctx, &theProject.Account, thePod, pod.ActionUpdate) {
		return responder.Response(http.StatusForbidden, "You do not have permission to update this pod"), nil
	}

	err = thePod.Update(ctx, podIn)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	lh.Debug(ctx, "Updated pod", "pod", thePod, "want", podIn)

	// TODO - Implement ProjectProjectIdPodPodIdPatch
	return responder.Response(http.StatusOK, thePod.ToAPI()), nil
}

// ProjectProjectIdPodPost - Creates a pod for a project.
func (s *PodAPIService) ProjectProjectIdPodPost(ctx context.Context, projectId string, accountID string, requestedPod api.Pod) (api.ImplResponse, error) {
	//lh.Debug(ctx, "Creating pod", "project_id", projectId, "acc", accountID, "pod", requestedPod)
	// Get the account by ID
	theProject, err := getProject(ctx, accountID, projectId)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	if !iam.RequestorCan(ctx, &theProject.Account, theProject, pod.ActionCreate) {
		return responder.Response(http.StatusForbidden, "You do not have permission to create this pod in this project"), nil
	}

	lh.Info(ctx, "Creating pod", "project_id", projectId, "acc", accountID, "pod", requestedPod)

	thepod, err := pod.Create(ctx, theProject, requestedPod)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}
	lh.Info(ctx, "Created pod", "pod", thepod)

	return responder.Response(http.StatusCreated, thepod.ToAPI()), nil
}

// ProjectProjectIdPodPodIdDelete - Deletes a pod for a project.
func (s *PodAPIService) ProjectProjectIdPodPodIdDelete(ctx context.Context, projectID string, podID string, accountID string) (api.ImplResponse, error) {

	log.Printf("%s %s %s", projectID, podID, accountID)

	theProject, err := getProject(ctx, accountID, projectID)
	if err != nil {
		return responder.Response(err.Code, err.Message), nil
	}

	// Get the pod by ID
	thePod, apiErr := pod.GetByID(ctx, theProject, podID)
	if apiErr != nil {
		return responder.Response(apiErr.Code, apiErr.Message), nil
	}

	// Check if the user can delete the pod
	if !iam.RequestorCan(ctx, &theProject.Account, thePod, pod.ActionDelete) {
		return responder.Response(http.StatusForbidden, "You do not have permission to delete this pod"), nil
	}

	// Delete the pod
	apiErr = thePod.Delete(ctx)
	if apiErr != nil {
		return responder.Response(apiErr.Code, apiErr.Message), nil
	}

	return responder.Response(http.StatusAccepted, nil), nil
}
