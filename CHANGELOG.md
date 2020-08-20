## 0.2.1 (Unreleased)
## 0.2.0 (July 06, 2020)

## 0.1.0 (Nov 22, 2019)

IMPROVEMENTS:

* CHANGELOG.md created
* Added GNUmakefile
* Added acceptance tests for `provider.go`, `data_source_rubrik_cluster_version.go`, `resource_rubrik_assign_sla.go`, `resource_rubrik_configure_timezone.go`
* Changed provider to look for upper case environment authentication variables
* Added check for lower case environment variables to compatibility with other Rubrik SDKs
* Converted existing code to utilize [Terraform plugin SDK](https://www.terraform.io/docs/extend/plugin-sdk.html)
* Added `go.mod` to support versioned [Go Modules](https://github.com/golang/go/wiki/Modules)