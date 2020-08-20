package rubrikcdm

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikAzureCloudOn() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikAzureCloudOnCreate,
		Read:   resourceRubrikAzureCloudOnRead,
		Update: resourceRubrikAzureCloudOnUpdate,
		Delete: resourceRubrikAzureCloudOnDelete,

		Schema: map[string]*schema.Schema{
			"archive_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the archive location used in the Rubrik GUI.",
			},
			"container": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Azure storage container being used as the archive target.",
			},
			"storage_account_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Storage Account that the container belongs to.",
			},
			"application_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The ID of the application registered in Azure Active Directory.",
			},
			"application_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The key of the application registered in Azure Active Directory.",
			},
			"directory_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The directory ID, also known as the tenant ID, found under the Azure Active Directory properties.",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"westus",
					"westus2",
					"centralus",
					"eastus",
					"eastus2",
					"northcentralus",
					"southcentralus",
					"westcentralus",
					"canadacentral",
					"canadaeast",
					"brazilsouth",
					"northeurope",
					"westeurope",
					"uksouth",
					"ukwest",
					"eastasia",
					"southeastasia",
					"japaneast",
					"japanwest",
					"australiaeast",
					"australiasoutheast",
					"centralindia",
					"southindia",
					"westindia",
					"koreacentral",
					"koreasouth",
				}, true),
				Description: "The name of the Azure region where the container is located.",
			},
			"virtual_network_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Azure virtual network ID used by Rubrik cluster to launch a temporary Rubrik instance in Azure for instantiation.",
			},
			"subnet_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Azure subnet name used by Rubrik cluster to launch a temporary Rubrik instance in Azure for instantiation.",
			},
			"security_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Azure Security Group ID used by Rubrik cluster to launch a temporary Rubrik instance in Azure for instantiation.",
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

func resourceRubrikAzureCloudOnCreate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	_, err := rubrik.AzureCloudOn(d.Get("archive_name").(string), d.Get("container").(string), d.Get("storage_account_name").(string), d.Get("application_id").(string), d.Get("application_key").(string), d.Get("directory_id").(string), d.Get("region").(string), d.Get("virtual_network_id").(string), d.Get("subnet_name").(string), d.Get("security_group_id").(string), d.Get("timeout").(int))
	if err != nil {

		if strings.Contains(err.Error(), "No change required") {
			d.SetId(d.Get("archive_name").(string))
			return resourceRubrikAzureCloudOnRead(d, meta)
		}

		return err
	}

	d.SetId(d.Get("archive_name").(string))

	return resourceRubrikAzureCloudOnRead(d, meta)
}

func resourceRubrikAzureCloudOnRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	var cloudOnConfigured = false
	archivesOnCluster, err := rubrik.CloudObjectStore(d.Get("timeout").(int))
	if err != nil {
		return err
	}

	for _, v := range archivesOnCluster.Data {

		if v.Definition.ObjectStoreType == "Azure" && v.Definition.Name == d.Get("archive_name").(string) {

			d.Set("archive_name", v.Definition.Name)

			d.Set("container", v.Definition.AzureComputeSummary.ContainerName)
			d.Set("storage_account_name", v.Definition.AzureComputeSummary.GeneralPurposeStorageAccountName)
			d.Set("application_id", v.Definition.AzureComputeSummary.ClientID)
			d.Set("directory_id", v.Definition.AzureComputeSummary.TenantID)
			d.Set("region", v.Definition.AzureComputeSummary.Region)

			d.Set("virtual_network_id", v.Definition.DefaultComputeNetworkConfig.VNetID)
			d.Set("subnet_name", v.Definition.DefaultComputeNetworkConfig.SubnetID)
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

func resourceRubrikAzureCloudOnUpdate(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	config := map[string]interface{}{}
	config["defaultComputeNetworkConfig"] = map[string]string{}
	config["azureComputeSummary"] = map[string]string{}

	config["defaultComputeNetworkConfig"].(map[string]string)["vNetId"] = d.Get("virtual_network_id").(string)
	config["defaultComputeNetworkConfig"].(map[string]string)["subnetId"] = d.Get("subnet_name").(string)
	config["defaultComputeNetworkConfig"].(map[string]string)["securityGroupId"] = d.Get("security_group_id").(string)

	config["azureComputeSummary"].(map[string]string)["subscriptionId"] = strings.Split(d.Get("virtual_network_id").(string), "/")[2]
	config["azureComputeSummary"].(map[string]string)["tenantId"] = d.Get("directory_id").(string)
	config["azureComputeSummary"].(map[string]string)["clientId"] = d.Get("application_id").(string)
	config["azureComputeSummary"].(map[string]string)["region"] = d.Get("region").(string)
	config["azureComputeSummary"].(map[string]string)["generalPurposeStorageAccountName"] = d.Get("storageAccountName").(string)
	config["azureComputeSummary"].(map[string]string)["containerName"] = d.Get("container").(string)
	config["azureComputeSummary"].(map[string]string)["clientId"] = d.Get("application_id").(string)
	config["azureComputeSummary"].(map[string]string)["securityGroupId"] = d.Get("security_group_id").(string)

	_, err := rubrik.UpdateCloudArchiveLocation(d.Get("archive_name").(string), config, d.Get("timeout").(int))
	if err != nil {
		if strings.Contains(err.Error(), "No change required") == true {
			return resourceRubrikAWSS3CloudOnRead(d, meta)
		}
		return err
	}

	return resourceRubrikAWSS3CloudOnRead(d, meta)
}

func resourceRubrikAzureCloudOnDelete(d *schema.ResourceData, meta interface{}) error {

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
