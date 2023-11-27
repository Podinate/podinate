/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Pod - A pod is a group of containers with the same lifecycle, and are the basic unit of deployment on Podinate
type Pod struct {

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

	// The global Resource ID of the pod
	ResourceId string `json:"resource_id,omitempty"`
}

// AssertPodRequired checks if the required fields are not zero-ed
func AssertPodRequired(obj Pod) error {
	return nil
}

// AssertRecursePodRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Pod (e.g. [][]Pod), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePodRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPod, ok := obj.(Pod)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPodRequired(aPod)
	})
}
