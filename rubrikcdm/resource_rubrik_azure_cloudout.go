package rubrikcdm

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAzureCloudOut() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAzureCloudOutCreate,
		Read:   resourceRubrikAzureCloudOutRead,
		Update: resourceRubrikAzureCloudOutUpdate,
		Delete: resourceRubrikAzureCloudOutDelete,

		Schema: map[string]*schema.Schema{
			"container": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Azure storage container you wish to use as an archive.",
			},
			"azure_access_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The access key for the Azure storage account.",
			},
			"storage_account_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Storage Account that the container belongs to.",
			},
			"archive_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the archive location used in the Rubrik GUI.",
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{
					"default",
					"china",
					"germany",
					"government",
				}, true),
				Description: "The Cloud Platform type of the archival location.",
			},
			"rsa_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The RSA key that will be used to encrypt the archive data.",
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

func resourceRubrikAzureCloudOutCreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.AzureCloudOut(d.Get("container").(string), d.Get("azure_access_key").(string), d.Get("storage_account_name").(string), d.Get("archive_name").(string), d.Get("instance_type").(string), d.Get("rsa_key").(string), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	d.SetId(d.Get("archive_name").(string))

	return resourceRubrikAzureCloudOutRead(d, meta)
}

func resourceRubrikAzureCloudOutRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	archivesOnCluster, err := rubrik.CloudObjectStore(d.Get("timeout").(int))
	if err != nil {
		return err
	}

	var archivePresent = false
	for _, v := range archivesOnCluster.Data {

		if v.Definition.ObjectStoreType == "Azure" && v.Definition.Name == d.Get("archive_name").(string) {
			d.Set("container", v.Definition.Bucket)
			d.Set("storage_account_name", v.Definition.AccessKey)
			d.Set("archive_name", v.Definition.Name)
			archivePresent = true
			break
		}
	}

	if archivePresent == false {
		d.SetId("")
	}

	return nil

}

func resourceRubrikAzureCloudOutUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := make(map[string]interface{})

	var archiveName string
	if d.HasChange("archive_name") {
		config["name"] = d.Get("archive_name").(string)
		old, _ := d.GetChange("archive_name")
		archiveName = old.(string)
	} else {
		archiveName = d.Get("archive_name").(string)
	}

	if d.HasChange("storage_account_name") {
		config["accessKey"] = d.Get("storage_account_name").(string)
	}

	if d.HasChange("azure_access_key") {
		config["secretKey"] = d.Get("azure_access_key").(string)
	}

	if len(config) == 0 {
		return resourceRubrikAzureCloudOutRead(d, meta)
	}

	_, err := rubrik.UpdateCloudArchiveLocation(archiveName, config, d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return err
		}
		return err
	}

	return resourceRubrikAzureCloudOutRead(d, meta)
}

func resourceRubrikAzureCloudOutDelete(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.RemoveArchiveLocation(d.Get("archive_name").(string), d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return nil
		}

		return err
	}
	return nil
}
