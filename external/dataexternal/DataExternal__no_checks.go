// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataexternal

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataExternal) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataExternal) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDataExternal_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDataExternal_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataExternal_IsTerraformDataSourceParameters(x interface{}) error {
	return nil
}

func validateDataExternal_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DataExternal) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataExternal) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DataExternal) validateSetProgramParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_DataExternal) validateSetQueryParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_DataExternal) validateSetWorkingDirParameters(val *string) error {
	return nil
}

func validateNewDataExternalParameters(scope constructs.Construct, id *string, config *DataExternalConfig) error {
	return nil
}

