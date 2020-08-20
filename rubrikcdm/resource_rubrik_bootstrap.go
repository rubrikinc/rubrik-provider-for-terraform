package rubrikcdm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func resourceRubrikBootstrap() *schema.Resource {
	return &schema.Resource{
		Create: resourceRubrikBootstrapCreate,
		Read:   resourceRubrikBootstrapRead,
		Update: resourceRubrikBootstrapUpdate,
		Delete: resourceRubrikBootstrapDelete,

		Schema: map[string]*schema.Schema{
			"cluster_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name to assign to the Rubrik cluster.",
			},
			"admin_email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Rubrik cluster sends messages for the admin account to this email address.",
			},
			"admin_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password for the admin account.",
			},
			"management_gateway": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.SingleIP(),
				Description:  "IP address assigned to the management network gateway",
			},
			"management_subnet_mask": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subnet mask assigned to the management network.",
			},
			"dns_search_domain": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The search domain that the DNS Service will use to resolve hostnames that are not fully qualified.",
			},
			"dns_name_servers": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "IPv4 addresses of DNS servers.",
			},
			"ntp_servers": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "FQDN or IPv4 address of a network time protocol (NTP) server.",
			},
			"node_config": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "The Node Name and IP formatted as a map.",
			},
			"enable_encryption": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable software data encryption at rest. When bootstraping a Cloud Cluster this value needs to be False.",
			},
			"wait_for_completion": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Flag to determine if the function should wait for the bootstrap process to complete.",
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

func resourceRubrikBootstrapCreate(d *schema.ResourceData, meta interface{}) error {

	// Convert interface{} list and maps to string
	dnsSearchDomain := make([]string, len(d.Get("dns_search_domain").([]interface{})))
	for i, v := range d.Get("dns_search_domain").([]interface{}) {
		dnsSearchDomain[i] = fmt.Sprint(v)
	}

	dnsNameServers := make([]string, len(d.Get("dns_name_servers").([]interface{})))
	for i, v := range d.Get("dns_name_servers").([]interface{}) {
		dnsNameServers[i] = fmt.Sprint(v)
	}

	ntpServers := make([]string, len(d.Get("ntp_servers").([]interface{})))
	for i, v := range d.Get("ntp_servers").([]interface{}) {
		ntpServers[i] = fmt.Sprint(v)
	}

	nodeConfig := make(map[string]string)
	for key, value := range d.Get("node_config").(map[string]interface{}) {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		nodeConfig[strKey] = strValue
	}

	rubrik := meta.(*rubrikcdm.Credentials)
	_, err := rubrik.Bootstrap(d.Get("cluster_name").(string), d.Get("admin_email").(string), d.Get("admin_password").(string), d.Get("management_gateway").(string), d.Get("management_subnet_mask").(string), dnsSearchDomain, dnsNameServers, ntpServers, nodeConfig, d.Get("enable_encryption").(bool), d.Get("wait_for_completion").(bool), d.Get("timeout").(int))
	if err != nil {
		return err
	}

	d.SetId(d.Get("cluster_name").(string))

	return resourceRubrikBootstrapRead(d, meta)
}

func resourceRubrikBootstrapRead(d *schema.ResourceData, meta interface{}) error {

	rubrik := meta.(*rubrikcdm.Credentials)

	bootstrapStatus, err := rubrik.ClusterBootstrapStatus()
	if err != nil {
		return err
	}

	if bootstrapStatus {
		d.Set("cluster_name", d.Get("cluster_name").(string))
		d.Set("admin_email", d.Get("admin_email").(string))
		d.Set("admin_password", d.Get("admin_password").(string))
		d.Set("management_gateway", d.Get("management_gateway").(string))
		d.Set("management_subnet_mask", d.Get("management_subnet_mask").(string))
		d.Set("dns_search_domain", d.Get("dns_search_domain").([]interface{}))
		d.Set("dns_name_servers", d.Get("dns_name_servers").([]interface{}))
		d.Set("ntp_servers", d.Get("ntp_servers").([]interface{}))
		d.Set("node_config", d.Get("node_config").(map[string]interface{}))
		d.Set("enable_encryption", d.Get("enable_encryption").(bool))
		d.Set("wait_for_completion", d.Get("wait_for_completion").(bool))
	} else {
		d.SetId("")
	}

	return nil

}

func resourceRubrikBootstrapUpdate(d *schema.ResourceData, meta interface{}) error {
	// Once a Cluster has been bootstrapped it can not be updated through the bootstrap resource

	return resourceRubrikBootstrapRead(d, meta)
}

func resourceRubrikBootstrapDelete(d *schema.ResourceData, meta interface{}) error {
	// Once a Cluster has been bootstrapped it can not be deleted.

	return nil
}
