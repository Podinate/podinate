/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate. Login should be performed over oauth from [auth.podinate.com](https://auth.podinate.com)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// ProjectApiService is a service that implements the logic for the ProjectApiServicer
// This service should implement the business logic for every endpoint for the ProjectApi API.
// Include any external packages or services that will be required by this service.
type ProjectApiService struct {
}

// NewProjectApiService creates a default api service
func NewProjectApiService() ProjectApiServicer {
	return &ProjectApiService{}
}

// ProjectGet - Returns a list of projects.
func (s *ProjectApiService) ProjectGet(ctx context.Context, account string, page int32, limit int32) (ImplResponse, error) {
	// TODO - update ProjectGet with the required logic for this service method.
	// Add api_project_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Project{}) or use other options such as http.Ok ...
	//return Response(200, Project{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectGet method not implemented")
}

// ProjectIdDelete - Delete an existing project
func (s *ProjectApiService) ProjectIdDelete(ctx context.Context, id string, account string) (ImplResponse, error) {
	// TODO - update ProjectIdDelete with the required logic for this service method.
	// Add api_project_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(202, Project{}) or use other options such as http.Ok ...
	//return Response(202, Project{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectIdDelete method not implemented")
}

// ProjectIdGet - Get an existing project given by ID
func (s *ProjectApiService) ProjectIdGet(ctx context.Context, id string, account string) (ImplResponse, error) {
	// TODO - update ProjectIdGet with the required logic for this service method.
	// Add api_project_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Project{}) or use other options such as http.Ok ...
	//return Response(200, Project{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectIdGet method not implemented")
}

// ProjectIdPatch - Update an existing project
func (s *ProjectApiService) ProjectIdPatch(ctx context.Context, id string, account string, project Project) (ImplResponse, error) {
	// TODO - update ProjectIdPatch with the required logic for this service method.
	// Add api_project_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Project{}) or use other options such as http.Ok ...
	//return Response(200, Project{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectIdPatch method not implemented")
}

// ProjectPost - Create a new project
func (s *ProjectApiService) ProjectPost(ctx context.Context, account string, project Project) (ImplResponse, error) {
	// TODO - update ProjectPost with the required logic for this service method.
	// Add api_project_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Project{}) or use other options such as http.Ok ...
	//return Response(201, Project{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectPost method not implemented")
}
