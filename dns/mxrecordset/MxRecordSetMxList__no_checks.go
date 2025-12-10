// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mxrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MxRecordSetMxList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MxRecordSetMxList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MxRecordSetMxList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MxRecordSetMxList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MxRecordSetMxList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MxRecordSetMxList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MxRecordSetMxList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMxRecordSetMxListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

