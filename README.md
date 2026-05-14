# CDKTN Go bindings for hashicorp/dns provider version 3.6.0

This repo builds and publishes the [Terraform dns provider](https://registry.terraform.io/providers/hashicorp/dns/3.6.0/docs) bindings for [CDK Terrain](https://cdktn.io).

## Go Package

The go package is generated into the [`github.com/cdktn-io/cdktn-provider-dns-go`](https://github.com/cdktn-io/cdktn-provider-dns-go) package.

`go get github.com/cdktn-io/cdktn-provider-dns-go/dns/<version>`

Where `<version>` is the version of the prebuilt provider you would like to use e.g. `v11`. The full module name can be found
within the [go.mod](https://github.com/cdktn-io/cdktn-provider-dns-go/blob/main/dns/go.mod#L1) file.

## Docs

Find auto-generated docs for this provider [here](https://github.com/cdktn-io/cdktn-provider-dns/blob/main/docs/API.go.md).


## Versioning

This project is explicitly not tracking the Terraform dns provider version 1:1. In fact, it always tracks `latest` of `~> 3.2` with every release. If there are scenarios where you explicitly have to pin your provider version, you can do so by [generating the provider constructs manually](https://cdktn.io/docs/concepts/providers#import-providers).

These are the upstream dependencies:

* [CDK Terrain](https://cdktn.io) - Last official release
* [Terraform dns provider](https://registry.terraform.io/providers/hashicorp/dns/3.6.0)
* [Terraform Engine](https://terraform.io)

If there are breaking changes (backward incompatible) in any of the above, the major version of this project will be bumped.

## Features / Issues / Bugs

Please report bugs and issues to the [CDK Terrain](https://cdktn.io) project:

* [Create bug report](https://github.com/open-constructs/cdk-terrain/issues)
* [Create feature request](https://github.com/open-constructs/cdk-terrain/issues)

## Contributing

### Projen

This is mostly based on [Projen](https://projen.io), which takes care of generating the entire repository.

### cdktn-provider-project based on Projen

There's a custom [project builder](https://github.com/cdktn-io/cdktn-provider-project) which encapsulate the common settings for all `cdktn` prebuilt providers.


### Repository Management

The repository is managed by [CDKTN Repository Manager](https://github.com/cdktn-io/cdktn-repository-manager/).
