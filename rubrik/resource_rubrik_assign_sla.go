package rubrikcdm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAssignSLA() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAssignSLACreate,
		Read:   resourceRubrikAssignSLARead,
		Update: resourceRubrikAssignSLAUpdate,
		Delete: resourceRubrikAssignSLADelete,

		Schema: map[string]*schema.Schema{
			"object_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     15,
				Description: "The name of the Rubrik object you wish to assign to an SLA Domain.",
			},
			"object_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"vmware",
					"ahv",
				}, true),
				Description: "The Rubrik object type you want to assign to the SLA Domain.",
			},
			"sla_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the SLA Domain you wish to assign an object to. To exclude the object from all SLA assignments use `do not protect` as the `sla_name`. To assign the selected object to the SLA of the next higher level object use `clear` as the `sla_name`.",
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

func resourceRubrikAssignSLACreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	tfID := fmt.Sprintf("%s-assigned-sla-%s", d.Get("object_type").(string), d.Get("object_name").(string))

	_, err := rubrik.AssignSLA(d.Get("object_name").(string), d.Get("object_type").(string), d.Get("sla_name").(string), d.Get("timeout").(int))
	if err != nil {

		if strings.Contains(err.Error(), "No change required") {
			d.SetId(tfID)
			return resourceRubrikAssignSLARead(d, meta)
		}

		return err
	}

	d.SetId(tfID)

	return resourceRubrikAssignSLARead(d, meta)
}

func resourceRubrikAssignSLARead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	var slaID string
	var err error
	switch d.Get("sla_name").(string) {
	case "do not protect":
		slaID = "UNPROTECTED"
	case "clear":
		slaID = "INHERIT"
	default:
		slaID, err = rubrik.ObjectID(d.Get("sla_name").(string), "sla", d.Get("timeout").(int))
		if err != nil {
			return err
		}
	}

	vmID, err := rubrik.ObjectID(d.Get("object_name").(string), d.Get("object_type").(string), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	var currentSLAName string
	switch d.Get("object_type").(string) {
	case "vmware":
		vmSummary, err := rubrik.Get("v1", fmt.Sprintf("/vmware/vm/%s", vmID), d.Get("timeout").(int))
		if err != nil {
			return err
		}

		switch slaID {
		case "INHERIT":
			currentSLAName = vmSummary.(map[string]interface{})["configuredSlaDomainName"].(string)
		default:
			currentSLAName = vmSummary.(map[string]interface{})["effectiveSlaDomainName"].(string)
		}
	case "ahv":
		vmSummary, err := rubrik.Get("internal", fmt.Sprintf("/nutanix/vm/%s", vmID), d.Get("timeout").(int))
		if err != nil {
			return err
		}

		switch slaID {
		case "INHERIT":
			currentSLAName = vmSummary.(map[string]interface{})["configuredSlaDomainName"].(string)
		default:
			currentSLAName = vmSummary.(map[string]interface{})["effectiveSlaDomainName"].(string)
		}

	}

	tfID := fmt.Sprintf("%s-%s-assigned-sla-%s", d.Get("object_type").(string), d.Get("object_name").(string), d.Get("sla_name").(string))

	d.SetId(tfID)

	if err = d.Set("sla_domain", currentSLAName); err != nil {
		return err
	}
	if err = d.Set("object_type", d.Get("object_type").(string)); err != nil {
		return err
	}
	if err = d.Set("object_name", d.Get("object_name").(string)); err != nil {
		return err
	}

	return nil

}

func resourceRubrikAssignSLAUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.AssignSLA(d.Get("object_name").(string), d.Get("object_type").(string), d.Get("sla_name").(string), d.Get("timeout").(int))
	if err != nil {

		if strings.Contains(err.Error(), "No change required") {
			return err
		}
		return err
	}

	return resourceRubrikAssignSLARead(d, meta)
}

func resourceRubrikAssignSLADelete(d *schema.ResourceData, meta interface{}) error {
	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.AssignSLA(d.Get("object_name").(string), d.Get("object_type").(string), "clear", d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") {
			return nil
		}

		return err
	}

	return nil
}
