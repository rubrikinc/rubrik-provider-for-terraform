package rubrikcdm

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAWSS3CloudOut() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAWSS3CloudOutCreate,
		Read:   resourceRubrikAWSS3CloudOutRead,
		Update: resourceRubrikAWSS3CloudOutUpdate,
		Delete: resourceRubrikAWSS3CloudOutDelete,

		Schema: map[string]*schema.Schema{
			"aws_bucket": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the AWS S3 bucket you wish to use as an archive target.",
			},
			"storage_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"standard",
					"standard_ia",
					"reduced_redundancy",
				}, true),
				Default:     "standard",
				Description: "The AWS storage class you wish to use.",
			},
			"archive_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the archive location used in the Rubrik GUI.",
			},
			"aws_region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_DEFAULT_REGION", nil),
				ValidateFunc: validation.StringInSlice([]string{
					"ap-south-1",
					"ap-northeast-3",
					"ap-northeast-2",
					"ap-southeast-1",
					"ap-southeast-2",
					"ap-northeast-1",
					"ca-central-1",
					"cn-north-1",
					"cn-northwest-1",
					"eu-central-1",
					"eu-west-1",
					"eu-west-2",
					"eu-west-3",
					"us-west-1",
					"us-east-1",
					"us-east-2",
					"us-west-2",
				}, true),
				Description: "The name of the AWS region where the bucket is located.",
			},
			"aws_access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", nil),
				Sensitive:   true,
				Description: "The access key of a AWS account with the required permissions.",
			},
			"aws_secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", nil),
				Sensitive:   true,
				Description: "The secret key of a AWS account with the required permissions.",
			},
			"rsa_key": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"kms_master_key_id"},
				Sensitive:     true,
				Description:   "The RSA key that will be used to encrypt the archive data.",
			},
			"kms_master_key_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"rsa_key"},
				Sensitive:     true,
				Description:   "The AWS KMS master key ID that will be used to encrypt the archive data.",
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

func resourceRubrikAWSS3CloudOutCreate(d *schema.ResourceData, meta interface{}) error {

	_, rsaOk := d.GetOk("rsa_key")
	_, kmsOk := d.GetOk("kms_master_key_id")

	if !rsaOk && !kmsOk {
		return errors.New("Either `rsa_key` or `kms_master_key_id` must be provided")
	}

	rubrik := meta.(*rubrikcdm.Credentials)

	if rsaOk {
		_, err := rubrik.AWSS3CloudOutRSA(d.Get("aws_bucket").(string), d.Get("storage_class").(string), d.Get("archive_name").(string), d.Get("aws_region").(string), d.Get("aws_access_key").(string), d.Get("aws_secret_key").(string), d.Get("rsa_key").(string), d.Get("timeout").(int))
		if err != nil {
			return err
		}
	} else {
		_, err := rubrik.AWSS3CloudOutKMS(d.Get("aws_bucket").(string), d.Get("storage_class").(string), d.Get("archive_name").(string), d.Get("aws_region").(string), d.Get("aws_access_key").(string), d.Get("aws_secret_key").(string), d.Get("kms_master_key_id").(string), d.Get("timeout").(int))
		if err != nil {
			return err
		}

	}

	d.SetId(d.Get("aws_bucket").(string))

	return resourceRubrikAWSS3CloudOutRead(d, meta)
}

func resourceRubrikAWSS3CloudOutRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	archivesOnCluster, err := rubrik.CloudObjectStore(d.Get("timeout").(int))
	if err != nil {
		return err
	}

	var archivePresent = false
	for _, v := range archivesOnCluster.Data {

		if v.Definition.ObjectStoreType == "S3" && v.Definition.Name == d.Get("archive_name").(string) {
			d.Set("archive_name", v.Definition.Name)
			d.Set("aws_bucket", v.Definition.Bucket)
			d.Set("storage_class", strings.ToLower(v.Definition.StorageClass))
			d.Set("aws_region", v.Definition.DefaultRegion)
			d.Set("aws_access_key", v.Definition.AccessKey)

			archivePresent = true
			break
		}
	}

	if archivePresent == false {
		d.SetId("")
	}

	return nil

}

func resourceRubrikAWSS3CloudOutUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := make(map[string]interface{})
	if d.HasChange("storage_class") {
		config["storageClass"] = strings.ToUpper(d.Get("storage_class").(string))
	}

	var archiveName string
	if d.HasChange("archive_name") {
		config["name"] = d.Get("archive_name").(string)
		old, _ := d.GetChange("archive_name")
		archiveName = old.(string)
	} else {
		archiveName = d.Get("archive_name").(string)
	}

	if d.HasChange("aws_access_key") {
		config["accessKey"] = d.Get("aws_access_key").(string)
	}

	if d.HasChange("aws_secret_key") {
		config["secretKey"] = d.Get("aws_secret_key").(string)
	}

	if len(config) == 0 {
		return resourceRubrikAWSS3CloudOutRead(d, meta)
	}

	_, err := rubrik.UpdateCloudArchiveLocation(archiveName, config, d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return err
		}
		return err
	}

	return resourceRubrikAWSS3CloudOutRead(d, meta)
}

func resourceRubrikAWSS3CloudOutDelete(d *schema.ResourceData, meta interface{}) error {

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
