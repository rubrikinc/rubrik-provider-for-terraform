package rubrikcdm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAWSExportEC2() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAWSExportEC2Create,
		Read:   resourceRubrikAWSExportEC2Read,
		Update: resourceRubrikAWSExportEC2Update,
		Delete: resourceRubrikAWSExportEC2Delete,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Instance ID of the AWS EC2 instance you wish to export.",
			},
			"export_instance_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name to assign to instance being launched.",
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"a1.medium",
					"a1.large",
					"a1.xlarge",
					"a1.2xlarge",
					"a1.4xlarge",
					"m4.large",
					"m4.xlarge",
					"m4.2xlarge",
					"m4.4xlarge",
					"m4.10xlarge",
					"m4.16xlarge",
					"m5.large",
					"m5.xlarge",
					"m5.2xlarge",
					"m5.4xlarge",
					"m5.12xlarge",
					"m5.24xlarge",
					"m5a.large",
					"m5a.xlarge",
					"m5a.2xlarge",
					"m5a.4xlarge",
					"m5a.12xlarge",
					"m5a.24xlarge",
					"m5d.large",
					"m5d.xlarge",
					"m5d.2xlarge",
					"m5d.4xlarge",
					"m5d.12xlarge",
					"m5d.24xlarge",
					"t2.nano",
					"t2.micro",
					"t2.small",
					"t2.medium",
					"t2.large",
					"t2.xlarge",
					"t2.2xlarge",
					"t3.nano",
					"t3.micro",
					"t3.small",
					"t3.medium",
					"t3.large",
					"t3.xlarge",
					"t3.2xlarge",
					"c4.large",
					"c4.xlarge",
					"c4.2xlarge",
					"c4.4xlarge",
					"c4.8xlarge",
					"c5.large",
					"c5.xlarge",
					"c5.2xlarge",
					"c5.4xlarge",
					"c5.9xlarge",
					"c5.18xlarge",
					"c5d.xlarge",
					"c5d.2xlarge",
					"c5d.4xlarge",
					"c5d.9xlarge",
					"c5d.18xlarge",
					"c5n.large",
					"c5n.xlarge",
					"c5n.2xlarge",
					"c5n.4xlarge",
					"c5n.9xlarge",
					"c5n.18xlarge",
					"r4.large",
					"r4.xlarge",
					"r4.2xlarge",
					"r4.4xlarge",
					"r4.8xlarge",
					"r4.16xlarge",
					"r5.large",
					"r5.xlarge",
					"r5.2xlarge",
					"r5.4xlarge",
					"r5.12xlarge",
					"r5.24xlarge",
					"r5a.large",
					"r5a.xlarge",
					"r5a.2xlarge",
					"r5a.4xlarge",
					"r5a.12xlarge",
					"r5a.24xlarge",
					"r5d.large",
					"r5d.xlarge",
					"r5d.2xlarge",
					"r5d.4xlarge",
					"r5d.12xlarge",
					"r5d.24xlarge",
					"x1.16xlarge",
					"x1.32xlarge",
					"x1e.xlarge",
					"x1e.2xlarge",
					"x1e.4xlarge",
					"x1e.8xlarge",
					"x1e.16xlarge",
					"x1e.32xlarge",
					"z1d.large",
					"z1d.xlarge",
					"z1d.2xlarge",
					"z1d.3xlarge",
					"z1d.6xlarge",
					"z1d.12xlarge",
					"d2.xlarge",
					"d2.2xlarge",
					"d2.4xlarge",
					"d2.8xlarge",
					"h1.2xlarge",
					"h1.4xlarge",
					"h1.8xlarge",
					"h1.16xlarge",
					"i3.large",
					"i3.xlarge",
					"i3.2xlarge",
					"i3.4xlarge",
					"i3.8xlarge",
					"i3.16xlarge",
					"f1.2xlarge",
					"f1.4xlarge",
					"f1.16xlarge",
					"g3s.xlarge",
					"g3.4xlarge",
					"g3.8xlarge",
					"g3.16xlarge",
					"p2.xlarge",
					"p2.8xlarge",
					"p2.16xlarge",
					"p3.2xlarge",
					"p3.8xlarge",
					"p3.16xlarge",
					"p3dn.24xlarge",
				}, true),
				Description: "The EC2 Instance Type to use for the instance being launched.",
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
			"subnet_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the subnet to assign to the instance being launched.",
			},
			"security_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the security group to assign to the instance being launched.",
			},
			"date_time": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The date and time of the EC2 Snapshot you wish to export formated as 'Month:Day:Year Hour:Minute AM/PM'. Ex. 04-09-2019 05:56 PM. You may also use 'latest' to export the last snapshot taken.",
			},
			"wait_for_completion": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Boolean flag to determine if the resource should wait for the export job to complete before continuing.",
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

func resourceRubrikAWSExportEC2Create(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.ExportEC2Instance(d.Get("instance_id").(string), d.Get("export_instance_name").(string), d.Get("instance_type").(string), d.Get("aws_region").(string), d.Get("subnet_id").(string), d.Get("security_group_id").(string), d.Get("date_time").(string), d.Get("wait_for_completion").(bool), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	d.SetId(d.Get("instance_id").(string))

	return resourceRubrikAWSExportEC2Read(d, meta)
}

func resourceRubrikAWSExportEC2Read(d *schema.ResourceData, meta interface{}) error {

	d.SetId("")

	return nil

}

func resourceRubrikAWSExportEC2Update(d *schema.ResourceData, meta interface{}) error {

	return resourceRubrikAWSExportEC2Read(d, meta)
}

func resourceRubrikAWSExportEC2Delete(d *schema.ResourceData, meta interface{}) error {
	// Once an EC2 instance has been exported it's lifecycle can not be managed by the Rubrik Cluster
	return nil
}
