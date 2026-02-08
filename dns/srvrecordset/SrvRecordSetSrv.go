// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package srvrecordset


type SrvRecordSetSrv struct {
	// The port for the service on the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#port SrvRecordSet#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The priority for the record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#priority SrvRecordSet#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The FQDN of the target, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#target SrvRecordSet#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The weight for the record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs/resources/srv_record_set#weight SrvRecordSet#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

