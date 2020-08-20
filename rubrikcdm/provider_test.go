package rubrikcdm

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"rubrik": testAccProvider,
	}
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("RUBRIK_CDM_USERNAME"); v == "" {
		t.Fatal("RUBRIK_CDM_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("RUBRIK_CDM_PASSWORD"); v == "" {
		t.Fatal("RUBRIK_CDM_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("RUBRIK_CDM_NODE_IP"); v == "" {
		t.Fatal("RUBRIK_CDM_NODE_IP must be set for acceptance tests")
	}
}

func testAccCheckEnvVariables(t *testing.T, variableNames []string) {
	for _, name := range variableNames {
		if v := os.Getenv(name); v == "" {
			t.Skipf("%s must be set for this acceptance test", name)
		}
	}
}
