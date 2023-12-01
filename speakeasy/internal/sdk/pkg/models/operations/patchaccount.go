// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/podinate/terraform-provider-podinate/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchAccountRequest struct {
	// The account to use for the request
	Account  string          `header:"style=simple,explode=false,name=account"`
	Account1 *shared.Account `request:"mediaType=application/json"`
}

func (o *PatchAccountRequest) GetAccount() string {
	if o == nil {
		return ""
	}
	return o.Account
}

func (o *PatchAccountRequest) GetAccount1() *shared.Account {
	if o == nil {
		return nil
	}
	return o.Account1
}

type PatchAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Account updated successfully
	Account *shared.Account
	// Request issued incorrectly, for example missing parameters or wrong endpoint
	Error *shared.Error
}

func (o *PatchAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchAccountResponse) GetAccount() *shared.Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *PatchAccountResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}