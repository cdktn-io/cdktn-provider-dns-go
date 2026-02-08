// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ptrrecord

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PtrRecordConfig struct {
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
	// The canonical name this record will point to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/ptr_record#ptr PtrRecord#ptr}
	Ptr *string `field:"required" json:"ptr" yaml:"ptr"`
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/ptr_record#zone PtrRecord#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The name of the record.
	//
	// The `zone` argument will be appended to this value to create the full record path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/ptr_record#name PtrRecord#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The TTL of the record. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/ptr_record#ttl PtrRecord#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

