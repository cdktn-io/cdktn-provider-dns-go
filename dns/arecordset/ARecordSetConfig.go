// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ARecordSetConfig struct {
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
	// The IPv4 addresses this record set will point to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/a_record_set#addresses ARecordSet#addresses}
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/a_record_set#zone ARecordSet#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/a_record_set#id ARecordSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the record set.
	//
	// The `zone` argument will be appended to this value to create the full record path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/a_record_set#name ARecordSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The TTL of the record set. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/a_record_set#ttl ARecordSet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

