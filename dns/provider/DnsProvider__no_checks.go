// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DnsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDnsProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDnsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDnsProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDnsProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsProvider) validateSetUpdateParameters(val interface{}) error {
	return nil
}

func validateNewDnsProviderParameters(scope constructs.Construct, id *string, config *DnsProviderConfig) error {
	return nil
}

