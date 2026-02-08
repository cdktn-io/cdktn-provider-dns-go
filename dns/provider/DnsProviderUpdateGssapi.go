// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type DnsProviderUpdateGssapi struct {
	// This or `password` is required if `username` is set, not supported on Windows.
	//
	// The path to a keytab file containing a key for `username`. Value can also be sourced from the DNS_UPDATE_KEYTAB environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#keytab DnsProvider#keytab}
	Keytab *string `field:"optional" json:"keytab" yaml:"keytab"`
	// This or `keytab` is required if `username` is set.
	//
	// The matching password for `username`. Value can also be sourced from the DNS_UPDATE_PASSWORD environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#password DnsProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The Kerberos realm or Active Directory domain. Value can also be sourced from the DNS_UPDATE_REALM environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#realm DnsProvider#realm}
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// The name of the user to authenticate as.
	//
	// If not set the current user session will be used. Value can also be sourced from the DNS_UPDATE_USERNAME environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#username DnsProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

