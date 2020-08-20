package rubrikcdm

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAWSS3CloudOn() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAWSS3CloudOnCreate,
		Read:   resourceRubrikAWSS3CloudOnRead,
		Update: resourceRubrikAWSS3CloudOnUpdate,
		Delete: resourceRubrikAWSS3CloudOnDelete,

		Schema: map[string]*schema.Schema{
			"archive_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the archive location used in the Rubrik GUI.",
			},
			"vpc_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AWS VPC ID used by Rubrik cluster to launch a temporary Rubrik instance in AWS for instantiation.",
			},
			"subnet_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AWS Subnet ID used by Rubrik cluster to launch a temporary Rubrik instance in AWS for instantiation.",
			},
			"security_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AWS Security Group ID used by Rubrik cluster to launch a temporary Rubrik instance in AWS for instantiation.",
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

func resourceRubrikAWSS3CloudOnCreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.AWSS3CloudOn(d.Get("archive_name").(string), d.Get("vpc_id").(string), d.Get("subnet_id").(string), d.Get("security_group_id").(string), d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") {
			d.SetId(d.Get("vpc_id").(string))
			return resourceRubrikAWSS3CloudOnRead(d, meta)
		}
	}

	d.SetId(d.Get("vpc_id").(string))

	return resourceRubrikAWSS3CloudOnRead(d, meta)
}

func resourceRubrikAWSS3CloudOnRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	var cloudOnConfigured = false
	archivesOnCluster, err := rubrik.CloudObjectStore(d.Get("timeout").(int))
	if err != nil {
		return err
	}

	for _, v := range archivesOnCluster.Data {

		if v.Definition.ObjectStoreType == "S3" && v.Definition.Name == d.Get("archive_name").(string) {

			d.Set("archive_name", v.Definition.Name)
			d.Set("vpc_id", v.Definition.DefaultComputeNetworkConfig.VNetID)
			d.Set("subnet_id", v.Definition.DefaultComputeNetworkConfig.SubnetID)
			d.Set("security_group_id", v.Definition.DefaultComputeNetworkConfig.SecurityGroupID)

			cloudOnConfigured = true
			break
		}
	}

	if cloudOnConfigured == false {
		d.SetId("")
	}

	return nil

}

func resourceRubrikAWSS3CloudOnUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := map[string]interface{}{}

	config["defaultComputeNetworkConfig"] = map[string]string{}
	config["defaultComputeNetworkConfig"].(map[string]string)["vNetId"] = d.Get("vpc_id").(string)
	config["defaultComputeNetworkConfig"].(map[string]string)["subnetId"] = d.Get("subnet_id").(string)
	config["defaultComputeNetworkConfig"].(map[string]string)["securityGroupId"] = d.Get("security_group_id").(string)

	_, err := rubrik.UpdateCloudArchiveLocation(d.Get("archive_name").(string), config, d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return resourceRubrikAWSS3CloudOnRead(d, meta)
		}
		return err
	}

	return resourceRubrikAWSS3CloudOnRead(d, meta)
}

func resourceRubrikAWSS3CloudOnDelete(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := map[string]interface{}{}
	config["isComputeEnabled"] = false

	_, err := rubrik.UpdateCloudArchiveLocation(d.Get("archive_name").(string), config, d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return resourceRubrikAWSS3CloudOnRead(d, meta)
		}
		return err
	}

	return resourceRubrikAWSS3CloudOnRead(d, meta)
}
