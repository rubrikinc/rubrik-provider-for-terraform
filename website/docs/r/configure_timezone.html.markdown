---
layout: "rubrik"
page_title: "Rubrik : configure_timezone"
description: |-
  Set the timezone on a Rubrik CDM cluster.
---

# assign_sla

Provides the ability to set the time zone that is used by the Rubrik cluster which uses the specified time zone for time values in the web UI, all reports, SLA Domain settings and all other time related operations.

## Example Usage

```hcl
resource "rubrik_configure_timezone" "example" {
  timezone = "America/Los_Angeles"
}
```

## Argument Reference

The following arguments are supported:

* `timezone` - (Required) The timezone used by the Rubrik cluster which uses the specified time zone for time values in the web UI, all reports, SLA Domain settings, and all other time related operations. Valid choices are  America/Anchorage, America/Araguaina, America/Barbados, America/Chicago, America/Denver, America/Los_Angeles America/Mexico_City, America/New_York, America/Noronha, America/Phoenix, America/Toronto, America/Vancouver, Asia/Bangkok, Asia/Dhaka, Asia/Dubai, Asia/Hong_Kong, Asia/Karachi,Asia/Kathmandu, Asia/Kolkata, Asia/Magadan, Asia/Singapore, Asia/Tokyo, Atlantic/Cape_Verde, Australia/Perth, Australia/Sydney, Europe/Amsterdam, Europe/Athens, Europe/London, Europe/Moscow, Pacific/Auckland, Pacific/Honolulu, Pacific/Midway, or UTC.
* `timeout` - (Optional) The number of seconds to wait to establish a connection the Rubrik cluster before returning a timeout error. Default is 15.
