/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate. Login should be performed over oauth from [auth.podinate.com](https://auth.podinate.com)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserGet200Response struct {

	// The user's ID
	Id string `json:"id,omitempty"`

	// The user's email address
	Email string `json:"email,omitempty"`
}

// AssertUserGet200ResponseRequired checks if the required fields are not zero-ed
func AssertUserGet200ResponseRequired(obj UserGet200Response) error {
	return nil
}

// AssertRecurseUserGet200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserGet200Response (e.g. [][]UserGet200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserGet200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserGet200Response, ok := obj.(UserGet200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserGet200ResponseRequired(aUserGet200Response)
	})
}
