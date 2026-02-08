// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datadnsmxrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDnsMxRecordSetMxList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataDnsMxRecordSetMxList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDnsMxRecordSetMxList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDnsMxRecordSetMxList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDnsMxRecordSetMxList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDnsMxRecordSetMxList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDnsMxRecordSetMxListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

