// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type DnsProviderUpdate struct {
	// gssapi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#gssapi DnsProvider#gssapi}
	Gssapi interface{} `field:"optional" json:"gssapi" yaml:"gssapi"`
	// Required if `key_name` is set.
	//
	// When using TSIG authentication, the algorithm to use for HMAC. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha256` or `hmac-sha512`. Value can also be sourced from the DNS_UPDATE_KEYALGORITHM environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#key_algorithm DnsProvider#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// The name of the TSIG key used to sign the DNS update messages.
	//
	// Value can also be sourced from the DNS_UPDATE_KEYNAME environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#key_name DnsProvider#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Required if `key_name` is set A Base64-encoded string containing the shared secret to be used for TSIG.
	//
	// Value can also be sourced from the DNS_UPDATE_KEYSECRET environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#key_secret DnsProvider#key_secret}
	KeySecret *string `field:"optional" json:"keySecret" yaml:"keySecret"`
	// The target UDP port on the server where updates are sent to.
	//
	// Defaults to `53`. Value can also be sourced from the DNS_UPDATE_PORT environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#port DnsProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Enable the Recursion Desired (RD) flag on DNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#recursive DnsProvider#recursive}
	Recursive interface{} `field:"optional" json:"recursive" yaml:"recursive"`
	// How many times to retry on connection timeout.
	//
	// Defaults to `3`. Value can also be sourced from the DNS_UPDATE_RETRIES environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#retries DnsProvider#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The hostname or IP address of the DNS server to send updates to.
	//
	// Value can also be sourced from the DNS_UPDATE_SERVER environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#server DnsProvider#server}
	Server *string `field:"optional" json:"server" yaml:"server"`
	// Timeout for DNS queries.
	//
	// Valid values are durations expressed as `500ms`, etc. or a plain number which is treated as whole seconds. Value can also be sourced from the DNS_UPDATE_TIMEOUT environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#timeout DnsProvider#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Transport to use for DNS queries.
	//
	// Valid values are `udp`, `udp4`, `udp6`, `tcp`, `tcp4`, or `tcp6`. Any UDP transport will retry automatically with the equivalent TCP transport in the event of a truncated response. Defaults to `udp`. Value can also be sourced from the DNS_UPDATE_TRANSPORT environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.5.0/docs#transport DnsProvider#transport}
	Transport *string `field:"optional" json:"transport" yaml:"transport"`
}

