/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProjectProjectIdPodGet200ResponseInner struct {

	// The short name (slug/url) of the pod
	Id string `json:"id,omitempty"`

	// The name of the pod
	Name string `json:"name,omitempty"`

	// The container image to run for this pod
	Image string `json:"image,omitempty"`

	// The image tag to run for this pod
	Tag string `json:"tag,omitempty"`

	// The current status of the pod
	Status string `json:"status,omitempty"`

	// The date and time the pod was created
	CreatedAt string `json:"created_at,omitempty"`
}

// AssertProjectProjectIdPodGet200ResponseInnerRequired checks if the required fields are not zero-ed
func AssertProjectProjectIdPodGet200ResponseInnerRequired(obj ProjectProjectIdPodGet200ResponseInner) error {
	return nil
}

// AssertRecurseProjectProjectIdPodGet200ResponseInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ProjectProjectIdPodGet200ResponseInner (e.g. [][]ProjectProjectIdPodGet200ResponseInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProjectProjectIdPodGet200ResponseInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProjectProjectIdPodGet200ResponseInner, ok := obj.(ProjectProjectIdPodGet200ResponseInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProjectProjectIdPodGet200ResponseInnerRequired(aProjectProjectIdPodGet200ResponseInner)
	})
}
