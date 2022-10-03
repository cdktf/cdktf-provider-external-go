//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExternalProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewExternalProviderParameters(scope constructs.Construct, id *string, config *ExternalProviderConfig) error {
	return nil
}

