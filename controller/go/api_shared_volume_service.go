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
	"context"
	"errors"
	"net/http"
)

// SharedVolumeApiService is a service that implements the logic for the SharedVolumeApiServicer
// This service should implement the business logic for every endpoint for the SharedVolumeApi API.
// Include any external packages or services that will be required by this service.
type SharedVolumeApiService struct {
}

// NewSharedVolumeApiService creates a default api service
func NewSharedVolumeApiService() SharedVolumeApiServicer {
	return &SharedVolumeApiService{}
}

// ProjectProjectIdSharedVolumesGet - Get a list of shared volumes for a given project
func (s *SharedVolumeApiService) ProjectProjectIdSharedVolumesGet(ctx context.Context, projectId string, account string, page int32, limit int32) (ImplResponse, error) {
	// TODO - update ProjectProjectIdSharedVolumesGet with the required logic for this service method.
	// Add api_shared_volume_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ProjectProjectIdSharedVolumesGet200Response{}) or use other options such as http.Ok ...
	//return Response(200, ProjectProjectIdSharedVolumesGet200Response{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectProjectIdSharedVolumesGet method not implemented")
}

// ProjectProjectIdSharedVolumesPost - Create a new shared volume
func (s *SharedVolumeApiService) ProjectProjectIdSharedVolumesPost(ctx context.Context, projectId string, account string, sharedVolume SharedVolume) (ImplResponse, error) {
	// TODO - update ProjectProjectIdSharedVolumesPost with the required logic for this service method.
	// Add api_shared_volume_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, SharedVolume{}) or use other options such as http.Ok ...
	//return Response(201, SharedVolume{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectProjectIdSharedVolumesPost method not implemented")
}

// ProjectProjectIdSharedVolumesVolumeIdDelete - Delete a shared volume
func (s *SharedVolumeApiService) ProjectProjectIdSharedVolumesVolumeIdDelete(ctx context.Context, projectId string, volumeId string, account string) (ImplResponse, error) {
	// TODO - update ProjectProjectIdSharedVolumesVolumeIdDelete with the required logic for this service method.
	// Add api_shared_volume_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(202, {}) or use other options such as http.Ok ...
	//return Response(202, nil),nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectProjectIdSharedVolumesVolumeIdDelete method not implemented")
}

// ProjectProjectIdSharedVolumesVolumeIdGet - Get a shared volume by ID
func (s *SharedVolumeApiService) ProjectProjectIdSharedVolumesVolumeIdGet(ctx context.Context, projectId string, volumeId string, account string) (ImplResponse, error) {
	// TODO - update ProjectProjectIdSharedVolumesVolumeIdGet with the required logic for this service method.
	// Add api_shared_volume_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, SharedVolume{}) or use other options such as http.Ok ...
	//return Response(200, SharedVolume{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectProjectIdSharedVolumesVolumeIdGet method not implemented")
}

// ProjectProjectIdSharedVolumesVolumeIdPut - Update a shared volume&#39;s spec
func (s *SharedVolumeApiService) ProjectProjectIdSharedVolumesVolumeIdPut(ctx context.Context, projectId string, volumeId string, account string, sharedVolume SharedVolume) (ImplResponse, error) {
	// TODO - update ProjectProjectIdSharedVolumesVolumeIdPut with the required logic for this service method.
	// Add api_shared_volume_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, SharedVolume{}) or use other options such as http.Ok ...
	//return Response(200, SharedVolume{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProjectProjectIdSharedVolumesVolumeIdPut method not implemented")
}