package rubrikcdm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikExample() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikExampleCreate,
		Read:   resourceRubrikExampleRead,
		Update: resourceRubrikExampleUpdate,
		Delete: resourceRubrikExampleDelete,

		Schema: map[string]*schema.Schema{
			"aws_account_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the AWS account you wish to protect. This is the name that will be displayed in the Rubrik UI",
			},
			"aws_access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", nil),
				Description: "The access key of a AWS account with the required permissions.",
			},
			"aws_secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", nil),
				Sensitive:   true,
				Description: "The secret key of a AWS account with the required permissions.",
			},
			"aws_regions": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of AWS regions to protect in this AWS account.",
			},
			"bolt_config": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Sensitive:   true,
				Elem:        &schema.Schema{Type: schema.TypeMap},
				Description: "List of dicts containing per region bolt network configs.",
			},
			"delete_snapshots": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "On destory, this will delete all EBS snapshots driven by Rubrik from your AWS account",
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

func resourceRubrikExampleCreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	awsRegionsString := make([]string, len(d.Get("aws_regions").([]interface{})))
	for i, v := range d.Get("aws_regions").([]interface{}) {
		awsRegionsString[i] = fmt.Sprint(v)
	}

	_, err := rubrik.AddAWSNativeAccount(d.Get("aws_account_name").(string), d.Get("aws_access_key").(string), d.Get("aws_secret_key").(string), awsRegionsString, d.Get("bolt_config").([]interface{}), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	d.SetId(d.Get("aws_account_name").(string))

	return resourceRubrikExampleRead(d, meta)
}

func resourceRubrikExampleRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	var accountPresent = true
	aws, err := rubrik.AWSAccountSummary(d.Get("aws_account_name").(string))
	if err != nil {
		if strings.Contains(err.Error(), "AWS Native Account was not found on the Rubrik cluster") == true {
			accountPresent = false
		} else {
			return err
		}

	}

	if accountPresent == false {
		d.SetId("")
	} else {
		d.Set("aws_account_name", aws.Name)
		d.Set("aws_access_key", aws.AccessKey)
		d.Set("aws_account_name", aws.Name)
		d.Set("aws_regions", aws.Regions)
		d.Set("bolt_config", aws.RegionalBoltNetworkConfigs)

	}

	return nil

}

func resourceRubrikExampleUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := make(map[string]interface{})
	var accountName string
	if d.HasChange("aws_account_name") {
		config["name"] = strings.ToUpper(d.Get("aws_account_name").(string))
		old, _ := d.GetChange("aws_account_name")
		accountName = old.(string)
	} else {
		accountName = d.Get("aws_account_name").(string)
	}

	if d.HasChange("aws_access_key") {
		config["accessKey"] = d.Get("aws_access_key").(string)
	}

	if d.HasChange("aws_secret_key") {
		config["secretKey"] = d.Get("aws_secret_key").(string)
	}

	if d.HasChange("aws_regions") {
		config["regions"] = d.Get("aws_regions").([]interface{})
	}

	if d.HasChange("bolt_config") {
		config["regionalBoltNetworkConfigs"] = d.Get("bolt_config").(map[string]interface{})
	}

	if len(config) == 0 {
		return resourceRubrikExampleRead(d, meta)
	}

	_, err := rubrik.UpdateAWSNativeAccount(accountName, config, d.Get("timeout").(int))
	if err != nil {
		return err
	}

	return resourceRubrikExampleRead(d, meta)
}

func resourceRubrikExampleDelete(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.RemoveAWSAccount(d.Get("aws_account_name").(string), d.Get("delete_snapshots").(bool))
	if err != nil {
		if strings.Contains(err.Error(), "AWS Native Account was not found on the Rubrik cluster") == true {
			return nil
		}
		return err

	}

	return nil
}
