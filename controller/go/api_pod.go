/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// PodApiController binds http requests to an api service and writes the service results to the http response
type PodApiController struct {
	service      PodApiServicer
	errorHandler ErrorHandler
}

// PodApiOption for how the controller is set up.
type PodApiOption func(*PodApiController)

// WithPodApiErrorHandler inject ErrorHandler into controller
func WithPodApiErrorHandler(h ErrorHandler) PodApiOption {
	return func(c *PodApiController) {
		c.errorHandler = h
	}
}

// NewPodApiController creates a default api controller
func NewPodApiController(s PodApiServicer, opts ...PodApiOption) Router {
	controller := &PodApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PodApiController
func (c *PodApiController) Routes() Routes {
	return Routes{
		{
			"ProjectProjectIdPodGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/pod",
			c.ProjectProjectIdPodGet,
		},
		{
			"ProjectProjectIdPodPodIdCopyGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/pod/{pod_id}/copy",
			c.ProjectProjectIdPodPodIdCopyGet,
		},
		{
			"ProjectProjectIdPodPodIdCopyPut",
			strings.ToUpper("Put"),
			"/v0/project/{project_id}/pod/{pod_id}/copy",
			c.ProjectProjectIdPodPodIdCopyPut,
		},
		{
			"ProjectProjectIdPodPodIdDelete",
			strings.ToUpper("Delete"),
			"/v0/project/{project_id}/pod/{pod_id}",
			c.ProjectProjectIdPodPodIdDelete,
		},
		{
			"ProjectProjectIdPodPodIdExecPost",
			strings.ToUpper("Post"),
			"/v0/project/{project_id}/pod/{pod_id}/exec",
			c.ProjectProjectIdPodPodIdExecPost,
		},
		{
			"ProjectProjectIdPodPodIdGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/pod/{pod_id}",
			c.ProjectProjectIdPodPodIdGet,
		},
		{
			"ProjectProjectIdPodPodIdLogsGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/pod/{pod_id}/logs",
			c.ProjectProjectIdPodPodIdLogsGet,
		},
		{
			"ProjectProjectIdPodPodIdPut",
			strings.ToUpper("Put"),
			"/v0/project/{project_id}/pod/{pod_id}",
			c.ProjectProjectIdPodPodIdPut,
		},
		{
			"ProjectProjectIdPodPost",
			strings.ToUpper("Post"),
			"/v0/project/{project_id}/pod",
			c.ProjectProjectIdPodPost,
		},
	}
}

// ProjectProjectIdPodGet - Get a list of pods for a given project
func (c *PodApiController) ProjectProjectIdPodGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	projectIdParam := params["project_id"]
	accountParam := r.Header.Get("account")
	pageParam, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ProjectProjectIdPodGet(r.Context(), projectIdParam, accountParam, pageParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdCopyGet - Copy a file out of the Pod
func (c *PodApiController) ProjectProjectIdPodPodIdCopyGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	pathParam := query.Get("path")
	result, err := c.service.ProjectProjectIdPodPodIdCopyGet(r.Context(), projectIdParam, podIdParam, accountParam, pathParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdCopyPut - Copy a file into a Pod
func (c *PodApiController) ProjectProjectIdPodPodIdCopyPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	pathParam := query.Get("path")
	bodyParam := &os.File{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ProjectProjectIdPodPodIdCopyPut(r.Context(), projectIdParam, podIdParam, accountParam, pathParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdDelete - Delete a pod
func (c *PodApiController) ProjectProjectIdPodPodIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	result, err := c.service.ProjectProjectIdPodPodIdDelete(r.Context(), projectIdParam, podIdParam, accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdExecPost - Execute a command in a pod
func (c *PodApiController) ProjectProjectIdPodPodIdExecPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	commandParam := strings.Split(query.Get("command"), ",")
	interactiveParam, err := parseBoolParameter(query.Get("interactive"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	ttyParam, err := parseBoolParameter(query.Get("tty"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	bodyParam := &os.File{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ProjectProjectIdPodPodIdExecPost(r.Context(), projectIdParam, podIdParam, accountParam, commandParam, interactiveParam, ttyParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdGet - Get a pod by ID
func (c *PodApiController) ProjectProjectIdPodPodIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	result, err := c.service.ProjectProjectIdPodPodIdGet(r.Context(), projectIdParam, podIdParam, accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdLogsGet - Get the logs for a pod
func (c *PodApiController) ProjectProjectIdPodPodIdLogsGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	linesParam, err := parseInt32Parameter(query.Get("lines"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	followParam, err := parseBoolParameter(query.Get("follow"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := c.service.ProjectProjectIdPodPodIdLogsGet(r.Context(), projectIdParam, podIdParam, accountParam, linesParam, followParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPodIdPut - Update a pod's spec
func (c *PodApiController) ProjectProjectIdPodPodIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	podIdParam := params["pod_id"]
	accountParam := r.Header.Get("account")
	podParam := Pod{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&podParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPodRequired(podParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ProjectProjectIdPodPodIdPut(r.Context(), projectIdParam, podIdParam, accountParam, podParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdPodPost - Create a new pod
func (c *PodApiController) ProjectProjectIdPodPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	accountParam := r.Header.Get("account")
	podParam := Pod{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&podParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPodRequired(podParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ProjectProjectIdPodPost(r.Context(), projectIdParam, accountParam, podParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
