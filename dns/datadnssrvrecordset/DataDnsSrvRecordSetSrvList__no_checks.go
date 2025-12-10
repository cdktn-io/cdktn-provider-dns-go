// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datadnssrvrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDnsSrvRecordSetSrvList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataDnsSrvRecordSetSrvList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDnsSrvRecordSetSrvList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDnsSrvRecordSetSrvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDnsSrvRecordSetSrvList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDnsSrvRecordSetSrvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDnsSrvRecordSetSrvListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

