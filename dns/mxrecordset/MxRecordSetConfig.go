// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mxrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MxRecordSetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/mx_record_set#zone MxRecordSet#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// mx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/mx_record_set#mx MxRecordSet#mx}
	Mx interface{} `field:"optional" json:"mx" yaml:"mx"`
	// The name of the record set.
	//
	// The `zone` argument will be appended to this value to create the full record path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/mx_record_set#name MxRecordSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The TTL of the record set. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/mx_record_set#ttl MxRecordSet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

