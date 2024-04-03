/*
 * Podinate API
 *
 * The API for the simple containerisation solution Podinate.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Service - A service makes a Pod available to other Pods or to the internet
type Service struct {

	// The hostname of the service
	Name string `json:"name"`

	// The port to expose the service on
	Port int32 `json:"port"`

	// The port to forward traffic to, if different from the port. Can be left blank if the same as the port.
	TargetPort int32 `json:"targetPort,omitempty"`

	// The protocol to use for the service. Either http, tcp or udp.
	Protocol string `json:"protocol"`

	// The domain name to use for the service. If left blank, the service will only be available internally. If set, the service will be available externally at the given domain name.
	DomainName string `json:"domain_name,omitempty"`
}

// AssertServiceRequired checks if the required fields are not zero-ed
func AssertServiceRequired(obj Service) error {
	elements := map[string]interface{}{
		"name":     obj.Name,
		"port":     obj.Port,
		"protocol": obj.Protocol,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseServiceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Service (e.g. [][]Service), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aService, ok := obj.(Service)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceRequired(aService)
	})
}