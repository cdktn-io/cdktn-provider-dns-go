// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package srvrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SrvRecordSetConfig struct {
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
	// The name of the record set.
	//
	// The `zone` argument will be appended to this value to create the full record path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#name SrvRecordSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#zone SrvRecordSet#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// srv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#srv SrvRecordSet#srv}
	Srv interface{} `field:"optional" json:"srv" yaml:"srv"`
	// The TTL of the record set. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#ttl SrvRecordSet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

