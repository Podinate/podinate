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

// UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service      UserApiServicer
	errorHandler ErrorHandler
}

// UserApiOption for how the controller is set up.
type UserApiOption func(*UserApiController)

// WithUserApiErrorHandler inject ErrorHandler into controller
func WithUserApiErrorHandler(h ErrorHandler) UserApiOption {
	return func(c *UserApiController) {
		c.errorHandler = h
	}
}

// NewUserApiController creates a default api controller
func NewUserApiController(s UserApiServicer, opts ...UserApiOption) Router {
	controller := &UserApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UserApiController
func (c *UserApiController) Routes() Routes {
	return Routes{
		{
			"UserGet",
			strings.ToUpper("Get"),
			"/v0/user",
			c.UserGet,
		},
		{
			"UserLoginCallbackProviderGet",
			strings.ToUpper("Get"),
			"/v0/user/login/callback/{provider}",
			c.UserLoginCallbackProviderGet,
		},
		{
			"UserLoginCompleteGet",
			strings.ToUpper("Get"),
			"/v0/user/login/complete",
			c.UserLoginCompleteGet,
		},
		{
			"UserLoginInitiateGet",
			strings.ToUpper("Get"),
			"/v0/user/login/initiate",
			c.UserLoginInitiateGet,
		},
		{
			"UserLoginPost",
			strings.ToUpper("Post"),
			"/v0/user/login",
			c.UserLoginPost,
		},
		{
			"UserLoginRedirectTokenGet",
			strings.ToUpper("Get"),
			"/v0/user/login/redirect/{token}",
			c.UserLoginRedirectTokenGet,
		},
	}
}

// UserGet - Get the current user
func (c *UserApiController) UserGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLoginCallbackProviderGet - User login callback URL for oauth providers
func (c *UserApiController) UserLoginCallbackProviderGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	providerParam := params["provider"]
	result, err := c.service.UserLoginCallbackProviderGet(r.Context(), providerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLoginCompleteGet - Complete a user login
func (c *UserApiController) UserLoginCompleteGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tokenParam := query.Get("token")
	clientParam := query.Get("client")
	result, err := c.service.UserLoginCompleteGet(r.Context(), tokenParam, clientParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLoginInitiateGet - Get a login URL for oauth login
func (c *UserApiController) UserLoginInitiateGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	providerParam := query.Get("provider")
	result, err := c.service.UserLoginInitiateGet(r.Context(), providerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLoginPost - Login to Podinate
func (c *UserApiController) UserLoginPost(w http.ResponseWriter, r *http.Request) {
	userLoginPostRequestParam := UserLoginPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userLoginPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserLoginPostRequestRequired(userLoginPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserLoginPost(r.Context(), userLoginPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLoginRedirectTokenGet - User login redirect URL to oauth providers
func (c *UserApiController) UserLoginRedirectTokenGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenParam := params["token"]
	result, err := c.service.UserLoginRedirectTokenGet(r.Context(), tokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
