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
	"net/http"
)



// AccountApiRouter defines the required methods for binding the api requests to a responses for the AccountApi
// The AccountApiRouter implementation should parse necessary information from the http request,
// pass the data to a AccountApiServicer to perform the required actions, then write the service results to the http response.
type AccountApiRouter interface { 
	AccountDelete(http.ResponseWriter, *http.Request)
	AccountGet(http.ResponseWriter, *http.Request)
	AccountPatch(http.ResponseWriter, *http.Request)
	AccountPost(http.ResponseWriter, *http.Request)
}
// PodApiRouter defines the required methods for binding the api requests to a responses for the PodApi
// The PodApiRouter implementation should parse necessary information from the http request,
// pass the data to a PodApiServicer to perform the required actions, then write the service results to the http response.
type PodApiRouter interface { 
	ProjectProjectIdPodGet(http.ResponseWriter, *http.Request)
	ProjectProjectIdPodPodIdDelete(http.ResponseWriter, *http.Request)
	ProjectProjectIdPodPodIdGet(http.ResponseWriter, *http.Request)
	ProjectProjectIdPodPodIdPatch(http.ResponseWriter, *http.Request)
	ProjectProjectIdPodPost(http.ResponseWriter, *http.Request)
}
// ProjectApiRouter defines the required methods for binding the api requests to a responses for the ProjectApi
// The ProjectApiRouter implementation should parse necessary information from the http request,
// pass the data to a ProjectApiServicer to perform the required actions, then write the service results to the http response.
type ProjectApiRouter interface { 
	ProjectGet(http.ResponseWriter, *http.Request)
	ProjectIdDelete(http.ResponseWriter, *http.Request)
	ProjectIdGet(http.ResponseWriter, *http.Request)
	ProjectIdPatch(http.ResponseWriter, *http.Request)
	ProjectPost(http.ResponseWriter, *http.Request)
}
// UserApiRouter defines the required methods for binding the api requests to a responses for the UserApi
// The UserApiRouter implementation should parse necessary information from the http request,
// pass the data to a UserApiServicer to perform the required actions, then write the service results to the http response.
type UserApiRouter interface { 
	UserGet(http.ResponseWriter, *http.Request)
	UserLoginCallbackProviderGet(http.ResponseWriter, *http.Request)
	UserLoginCompleteGet(http.ResponseWriter, *http.Request)
	UserLoginInitiateGet(http.ResponseWriter, *http.Request)
	UserLoginRedirectTokenGet(http.ResponseWriter, *http.Request)
}


// AccountApiServicer defines the api actions for the AccountApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AccountApiServicer interface { 
	AccountDelete(context.Context, string) (ImplResponse, error)
	AccountGet(context.Context, string) (ImplResponse, error)
	AccountPatch(context.Context, string, Account) (ImplResponse, error)
	AccountPost(context.Context, Account) (ImplResponse, error)
}


// PodApiServicer defines the api actions for the PodApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PodApiServicer interface { 
	ProjectProjectIdPodGet(context.Context, string, string, int32, int32) (ImplResponse, error)
	ProjectProjectIdPodPodIdDelete(context.Context, string, string, string) (ImplResponse, error)
	ProjectProjectIdPodPodIdGet(context.Context, string, string, string) (ImplResponse, error)
	ProjectProjectIdPodPodIdPatch(context.Context, string, string, string, Pod) (ImplResponse, error)
	ProjectProjectIdPodPost(context.Context, string, string, Pod) (ImplResponse, error)
}


// ProjectApiServicer defines the api actions for the ProjectApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ProjectApiServicer interface { 
	ProjectGet(context.Context, string, int32, int32) (ImplResponse, error)
	ProjectIdDelete(context.Context, string, string) (ImplResponse, error)
	ProjectIdGet(context.Context, string, string) (ImplResponse, error)
	ProjectIdPatch(context.Context, string, string, Project) (ImplResponse, error)
	ProjectPost(context.Context, string, Project) (ImplResponse, error)
}


// UserApiServicer defines the api actions for the UserApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserApiServicer interface { 
	UserGet(context.Context) (ImplResponse, error)
	UserLoginCallbackProviderGet(context.Context, string) (ImplResponse, error)
	UserLoginCompleteGet(context.Context, string, string) (ImplResponse, error)
	UserLoginInitiateGet(context.Context, string) (ImplResponse, error)
	UserLoginRedirectTokenGet(context.Context, string) (ImplResponse, error)
}
