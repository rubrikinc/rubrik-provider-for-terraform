package rubrikcdm

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

var currentTz string

func TestAccRubrikConfigureTimezone_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCheckRubrikConfigureTimezonePreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRubrikTimezoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRubrikConfigureTimezoneConfig("UTC"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rubrik_configure_timezone.timezone", "timezone", "UTC"),
				),
			},
		},
	})
}

func getTimezone() (string, error) {
	timeout := 15
	rubrik := testAccProvider.Meta().(*rubrikcdm.Credentials)

	clusterSummary, err := rubrik.Get("v1", "/cluster/me", timeout)
	if err != nil {

		return "", err
	}

	currentTimezone := fmt.Sprintf("%v", clusterSummary.(map[string]interface{})["timezone"].(map[string]interface{})["timezone"])

	return currentTimezone, nil
}

func setTimezone(timezone string) error {
	timeout := 15
	rubrik := testAccProvider.Meta().(*rubrikcdm.Credentials)

	_, err := rubrik.ConfigureTimezone(timezone, timeout)
	if err != nil {
		return err
	}

	return nil
}

func testAccRubrikConfigureTimezoneConfig(timezone string) string {
	return fmt.Sprintf(`
resource "rubrik_configure_timezone" "timezone" {
	timezone = "%s"
	timeout = 15
	}
`, timezone)
}

func testAccCheckRubrikConfigureTimezonePreCheck(t *testing.T) {
	var err error

	log.Printf("Running Pre-check")
	currentTz, err = getTimezone()
	if err != nil {
		t.Fatal(err.Error())
	}

	if currentTz == "UTC" {
		t.Fatal("CDM Timezone is already set to UTC. Please set to another timezone before testing.")
	}
}

func testAccCheckRubrikTimezoneDestroy(s *terraform.State) error {
	log.Printf("Running Cleanup. Setting timezone back to %s", currentTz)
	err := setTimezone(currentTz)
	if err != nil {
		return err
	}

	return nil
}
