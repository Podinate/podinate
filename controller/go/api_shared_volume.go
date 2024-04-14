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
	"strings"

	"github.com/gorilla/mux"
)

// SharedVolumeApiController binds http requests to an api service and writes the service results to the http response
type SharedVolumeApiController struct {
	service      SharedVolumeApiServicer
	errorHandler ErrorHandler
}

// SharedVolumeApiOption for how the controller is set up.
type SharedVolumeApiOption func(*SharedVolumeApiController)

// WithSharedVolumeApiErrorHandler inject ErrorHandler into controller
func WithSharedVolumeApiErrorHandler(h ErrorHandler) SharedVolumeApiOption {
	return func(c *SharedVolumeApiController) {
		c.errorHandler = h
	}
}

// NewSharedVolumeApiController creates a default api controller
func NewSharedVolumeApiController(s SharedVolumeApiServicer, opts ...SharedVolumeApiOption) Router {
	controller := &SharedVolumeApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SharedVolumeApiController
func (c *SharedVolumeApiController) Routes() Routes {
	return Routes{
		{
			"ProjectProjectIdSharedVolumesGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/shared_volumes",
			c.ProjectProjectIdSharedVolumesGet,
		},
		{
			"ProjectProjectIdSharedVolumesPost",
			strings.ToUpper("Post"),
			"/v0/project/{project_id}/shared_volumes",
			c.ProjectProjectIdSharedVolumesPost,
		},
		{
			"ProjectProjectIdSharedVolumesVolumeIdDelete",
			strings.ToUpper("Delete"),
			"/v0/project/{project_id}/shared_volumes/{volume_id}",
			c.ProjectProjectIdSharedVolumesVolumeIdDelete,
		},
		{
			"ProjectProjectIdSharedVolumesVolumeIdGet",
			strings.ToUpper("Get"),
			"/v0/project/{project_id}/shared_volumes/{volume_id}",
			c.ProjectProjectIdSharedVolumesVolumeIdGet,
		},
		{
			"ProjectProjectIdSharedVolumesVolumeIdPut",
			strings.ToUpper("Put"),
			"/v0/project/{project_id}/shared_volumes/{volume_id}",
			c.ProjectProjectIdSharedVolumesVolumeIdPut,
		},
	}
}

// ProjectProjectIdSharedVolumesGet - Get a list of shared volumes for a given project
func (c *SharedVolumeApiController) ProjectProjectIdSharedVolumesGet(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ProjectProjectIdSharedVolumesGet(r.Context(), projectIdParam, accountParam, pageParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdSharedVolumesPost - Create a new shared volume
func (c *SharedVolumeApiController) ProjectProjectIdSharedVolumesPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	accountParam := r.Header.Get("account")
	sharedVolumeParam := SharedVolume{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&sharedVolumeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSharedVolumeRequired(sharedVolumeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ProjectProjectIdSharedVolumesPost(r.Context(), projectIdParam, accountParam, sharedVolumeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdSharedVolumesVolumeIdDelete - Delete a shared volume
func (c *SharedVolumeApiController) ProjectProjectIdSharedVolumesVolumeIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	volumeIdParam := params["volume_id"]
	accountParam := r.Header.Get("account")
	result, err := c.service.ProjectProjectIdSharedVolumesVolumeIdDelete(r.Context(), projectIdParam, volumeIdParam, accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdSharedVolumesVolumeIdGet - Get a shared volume by ID
func (c *SharedVolumeApiController) ProjectProjectIdSharedVolumesVolumeIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	volumeIdParam := params["volume_id"]
	accountParam := r.Header.Get("account")
	result, err := c.service.ProjectProjectIdSharedVolumesVolumeIdGet(r.Context(), projectIdParam, volumeIdParam, accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ProjectProjectIdSharedVolumesVolumeIdPut - Update a shared volume's spec
func (c *SharedVolumeApiController) ProjectProjectIdSharedVolumesVolumeIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["project_id"]
	volumeIdParam := params["volume_id"]
	accountParam := r.Header.Get("account")
	sharedVolumeParam := SharedVolume{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&sharedVolumeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSharedVolumeRequired(sharedVolumeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ProjectProjectIdSharedVolumesVolumeIdPut(r.Context(), projectIdParam, volumeIdParam, accountParam, sharedVolumeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
