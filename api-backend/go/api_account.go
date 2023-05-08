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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// AccountApiController binds http requests to an api service and writes the service results to the http response
type AccountApiController struct {
	service AccountApiServicer
	errorHandler ErrorHandler
}

// AccountApiOption for how the controller is set up.
type AccountApiOption func(*AccountApiController)

// WithAccountApiErrorHandler inject ErrorHandler into controller
func WithAccountApiErrorHandler(h ErrorHandler) AccountApiOption {
	return func(c *AccountApiController) {
		c.errorHandler = h
	}
}

// NewAccountApiController creates a default api controller
func NewAccountApiController(s AccountApiServicer, opts ...AccountApiOption) Router {
	controller := &AccountApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AccountApiController
func (c *AccountApiController) Routes() Routes {
	return Routes{ 
		{
			"AccountDelete",
			strings.ToUpper("Delete"),
			"/v0/account",
			c.AccountDelete,
		},
		{
			"AccountGet",
			strings.ToUpper("Get"),
			"/v0/account",
			c.AccountGet,
		},
		{
			"AccountPatch",
			strings.ToUpper("Patch"),
			"/v0/account",
			c.AccountPatch,
		},
		{
			"AccountPost",
			strings.ToUpper("Post"),
			"/v0/account",
			c.AccountPost,
		},
	}
}

// AccountDelete - Delete the account and all associated resources!
func (c *AccountApiController) AccountDelete(w http.ResponseWriter, r *http.Request) {
	accountParam := r.Header.Get("account")
	result, err := c.service.AccountDelete(r.Context(), accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AccountGet - Get information about the current account.
func (c *AccountApiController) AccountGet(w http.ResponseWriter, r *http.Request) {
	accountParam := r.Header.Get("account")
	result, err := c.service.AccountGet(r.Context(), accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AccountPatch - Update an existing account
func (c *AccountApiController) AccountPatch(w http.ResponseWriter, r *http.Request) {
	accountParam := r.Header.Get("account")
	account2Param := Account{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&account2Param); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAccountRequired(account2Param); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AccountPatch(r.Context(), accountParam, account2Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AccountPost - Create a new account
func (c *AccountApiController) AccountPost(w http.ResponseWriter, r *http.Request) {
	accountParam := Account{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&accountParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAccountRequired(accountParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AccountPost(r.Context(), accountParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
