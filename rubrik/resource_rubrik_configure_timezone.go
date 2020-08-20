package rubrikcdm

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikConfigureTimezone() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikConfigureTimezoneCreate,
		Read:   resourceRubrikConfigureTimezoneRead,
		Update: resourceRubrikConfigureTimezoneUpdate,
		Delete: resourceRubrikConfigureTimezoneDelete,

		Schema: map[string]*schema.Schema{
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"America/Anchorage",
					"America/Araguaina",
					"America/Barbados",
					"America/Chicago",
					"America/Denver",
					"America/Los_Angeles",
					"America/Mexico_City",
					"America/New_York",
					"America/Noronha",
					"America/Phoenix",
					"America/Toronto",
					"America/Vancouver",
					"Asia/Bangkok",
					"Asia/Dhaka",
					"Asia/Dubai",
					"Asia/Hong_Kong",
					"Asia/Karachi",
					"Asia/Kathmandu",
					"Asia/Kolkata",
					"Asia/Magadan",
					"Asia/Singapore",
					"Asia/Tokyo",
					"Atlantic/Cape_Verde",
					"Australia/Perth",
					"Australia/Sydney",
					"Europe/Amsterdam",
					"Europe/Athens",
					"Europe/London",
					"Europe/Moscow",
					"Pacific/Auckland",
					"Pacific/Honolulu",
					"Pacific/Midway",
					"UTC",
				}, true),
				Description: "The timezone used by the Rubrik cluster which uses the specified time zone for time values in the web UI, all reports, SLA Domain settings, and all other time related operations",
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     15,
				Description: "The number of seconds to wait to establish a connection the Rubrik cluster before returning a timeout error.",
			},
		},
	}

}

func resourceRubrikConfigureTimezoneCreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.ConfigureTimezone(d.Get("timezone").(string), d.Get("timeout").(int))
	if err != nil {

		if strings.Contains(err.Error(), "No change required") {
			d.SetId("rubrik-cluster-timezone")
			return resourceRubrikConfigureTimezoneRead(d, meta)
		}

		return err
	}

	d.SetId("rubrik-cluster-timezone")

	return resourceRubrikConfigureTimezoneRead(d, meta)
}

func resourceRubrikConfigureTimezoneRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	log.Println("[INFO] Determining the current Rubrik cluster timezone.")
	clusterSummary, err := rubrik.Get("v1", "/cluster/me", d.Get("timeout").(int))
	if err != nil {

		return err
	}

	currentTimezone := clusterSummary.(map[string]interface{})["timezone"].(map[string]interface{})["timezone"]

	d.SetId("rubrik-cluster-timezone")

	err = d.Set("timezone", currentTimezone)
	if err != nil {
		return err
	}

	return nil

}

func resourceRubrikConfigureTimezoneUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.ConfigureTimezone(d.Get("timezone").(string), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	return nil
}

func resourceRubrikConfigureTimezoneDelete(d *schema.ResourceData, m interface{}) error {
	// Cluster Timezone is a requirement for the Rubrik cluster and can not be "deleted"
	return nil
}
