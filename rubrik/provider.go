package rubrikcdm

import (
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	// Look for environment variables from other Rubrik SDKs, and use them if necessary
	if os.Getenv("RUBRIK_CDM_NODE_IP") == "" && os.Getenv("rubrik_cdm_node_ip") != "" {
		os.Setenv("RUBRIK_CDM_NODE_IP", os.Getenv("rubrik_cdm_node_ip"))
		log.Printf("Setting environment variable RUBRIK_CDM_NODE_IP to match rubrik_cdm_node_ip")
	}

	if os.Getenv("RUBRIK_CDM_USERNAME") == "" && os.Getenv("rubrik_cdm_username") != "" {
		os.Setenv("RUBRIK_CDM_USERNAME", os.Getenv("rubrik_cdm_username"))
		log.Printf("Setting environment variable RUBRIK_CDM_USERNAME to match rubrik_cdm_username")
	}

	if os.Getenv("RUBRIK_CDM_PASSWORD") == "" && os.Getenv("rubrik_cdm_password") != "" {
		os.Setenv("RUBRIK_CDM_PASSWORD", os.Getenv("rubrik_cdm_password"))
		log.Printf("Setting environment variable RUBRIK_CDM_PASSWORD to match rubrik_cdm_password")
	}

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"node_ip": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("RUBRIK_CDM_NODE_IP", nil),
				ValidateFunc: validation.SingleIP(),
				Description:  "The IP Address of a Node in the Rubrik cluster.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RUBRIK_CDM_USERNAME", nil),
				Description: "The username used to authenticate against the Rubrik cluster.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RUBRIK_CDM_PASSWORD", nil),
				Description: "The password used to authenticate against the Rubrik cluster.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rubrik_assign_sla":         resourceRubrikAssignSLA(),
			"rubrik_configure_timezone": resourceRubrikConfigureTimezone(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"rubrik_cluster_version": dataSourceRubrikClusterVersion(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		NodeIP:   d.Get("node_ip").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return config.Client()
}
