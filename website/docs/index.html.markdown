---
layout: "rubrik"
page_title: "Provider: rubrik"
description: |-
  The Rubrik provider is used to interact with resources on the Rubrik CDM platform.
---

# Rubrik Provider

The Rubrik provider is used to interact with resources on the Rubrik CDM platform.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Rubrik Provider
provider "rubrik" {
  node_ip     = "${var.rubrik_node_ip}"
  username    = "${var.rubrik_username}"
  password    = "${var.rubrik_password}"
}

# Set the Rubrik cluster timezone
resource "rubrik_configure_timezone" "LA-Timezone" {
  timezone = "America/Los_Angeles"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `node_ip` - (Optional) The Node IP address of the Rubrik cluster
  you wish to connect to. The value may also be sourced from the `rubrik_node_ip`
  environment variable.

* `username` - (Optional) The username of the Rubrik cluster
  you wish to connect to. The value may also be sourced from the `rubrik_username`
  environment variable.

* `password` - (Optional) The password of the Rubrik cluster
  you wish to connect to. The value may also be sourced from the `rubrik_password`
  environment variable
