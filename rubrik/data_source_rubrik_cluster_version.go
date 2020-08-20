package rubrikcdm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func dataSourceRubrikClusterVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRubrikClusterVersionRead,

		Schema: map[string]*schema.Schema{
			"cluster_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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

func dataSourceRubrikClusterVersionRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*rubrikcdm.Credentials)

	clusterVersion, err := client.ClusterVersion(d.Get("timeout").(int))
	if err != nil {
		return err
	}

	d.SetId(clusterVersion)
	err = d.Set("cluster_version", clusterVersion)
	if err != nil {
		return err
	}

	return nil
}
