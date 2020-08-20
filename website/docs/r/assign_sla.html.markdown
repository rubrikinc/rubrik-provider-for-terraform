---
layout: "rubrik"
page_title: "Rubrik : assign_sla"
description: |-
  Assign a Rubrik object to a specified SLA Domain.
---

# assign_sla

Assign a Rubrik object to a specified SLA Domain.

## Example Usage

```hcl
resource "rubrik_assign_sla" "assign-sla" {
  object_name = "tf-example-vm"
  object_type = "vmware"
  sla_name    = "Gold"
}
```

## Argument Reference

The following arguments are supported:

* `object_name` - (Required) The name of the Rubrik object you wish to assign to an SLA Domain.
* `object_type` - (Required) The Rubrik object type you want to assign to the SLA Domain. Currently, `vmware` and `ahv` are the only supported values.
* `sla_name` - (Required) The name of the SLA Domain you wish to assign an object to. To exclude the object from all SLA assignments use `do not protect` as the `sla_name`. To assign the selected object to the SLA of the next higher level object use `clear` as the `sla_name`.
* `timeout` - (Optional) The number of seconds to wait to establish a connection the Rubrik cluster before returning a timeout error. Default is 15.

Valid object_type choices:
* vmware
* ahv

## Attribute Reference

The following attributes are exported:

* `id` - An ID unique to Terraform for this port group. The convention is a prefix, the host system ID, and the port group name. An example would be `vmware-exampleVMName-assigned-sla-Gold`.
* `sla_domain` - The name of the SLA Domain assigned to the `object_name`.
* `object_type` -  The type of the Rubrik object you assigned an SLA Domain to.
* `object_name` - The name of the Rubrik object you assigned an SLA Domain to.

## Destroy Behavior

On `terraform destroy`, this resource will set the `sla_domain` to `clear` which will assign the select `object_name` to the SLA of the next higher level object.
