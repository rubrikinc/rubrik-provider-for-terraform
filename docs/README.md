# Rubrik Provider

The Rubrik Provider transforms the Rubrik RESTful API functionality into easy to consume Terraform configuration whichs eliminates the need to understand how to consume raw Rubrik APIs extends upon one of Rubrik's main design centers - simplicity

## Example Usage

```hcl
provider "rubrik" {}

resource "rubrik_configure_timezone" "LA-Timezone" {
  timezone = "America/Los_Angeles"
}
```

## Authentication

The Rubrik provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials 

Static credentials can be provided by adding an `node_ip`, `username` and `password` in-line in the
Rubrik provider block:

Usage:

```hcl
provider "rubrik" {
  node_ip     = "10.255.41.201"
  username    = "admin"
  password    = "RubrikTFDemo2019"
}
```
### Environment variables

You can provide your credentials via the `RUBRIK_CDM_NODE_IP`, `RUBRIK_CDM_USERNAME` and
`RUBRIK_CDM_PASSWORD`, environment variables, representing your Rubrik Node IP address, username
and password, respectively.

```hcl
provider "rubrik" {}
```

```sh
$ export RUBRIK_CDM_NODE_IP="10.255.41.201"
$ export RUBRIK_CDM_USERNAME="admin"
$ export RUBRIK_CDM_PASSWORD="RubrikTFDemo2019"
$ terraform plan
```

## Argument Reference

The following arguments are supported in the Rubrik `provider` block:

* `node_ip` - (Optional) The Node IP address of the Rubrik cluster you wish to connect to. The value may also be sourced from the
`RUBRIK_CDM_PASSWORD` environment variable.

* `username` - (Optional) The username of the Rubrik cluster you wish to connect to. The value may also be sourced from the
`RUBRIK_CDM_USERNAME` environment variable.

* `password` - (Optional) The password of the Rubrik cluster you wish to connect to. The value may also be sourced from the
`RUBRIK_CDM_PASSWORD` environment variable.
