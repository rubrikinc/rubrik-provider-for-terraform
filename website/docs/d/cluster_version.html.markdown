---
layout: "rubrik"
page_title: "Rubrik : cluster_version"
description: |-
  Get the CDM version of the Rubrik cluster.
---

# cluster_version

Use this data source to get the CDM version of the Rubrik cluster.

## Example Usage

```hcl
data "rubrik_cluster_version" "version" {}
```

## Arguments Reference

None
 
## Attributes Reference

The following computed attributes are exported:

* cluster_version - The CDM version of the Rubrik cluster.
