// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExternalProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateExternalProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateExternalProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewExternalProviderParameters(scope constructs.Construct, id *string, config *ExternalProviderConfig) error {
	return nil
}

