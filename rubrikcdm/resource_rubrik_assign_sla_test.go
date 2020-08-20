package rubrikcdm

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

func TestAccRubrikAssignSla_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCheckEnvVariables(t, []string{"RUBRIK_CDM_TEST_VM", "RUBRIK_CDM_TEST_SLA"})
			testAccCheckRubrikAssignSLAPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRubrikAssignSLADestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRubrikAssignSLAConfig(os.Getenv("RUBRIK_CDM_TEST_VM"), os.Getenv("RUBRIK_CDM_TEST_SLA")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("rubrik_assign_sla.assign_sla", "object_name", os.Getenv("RUBRIK_CDM_TEST_VM")),
					resource.TestCheckResourceAttr("rubrik_assign_sla.assign_sla", "sla_name", os.Getenv("RUBRIK_CDM_TEST_SLA")),
				),
			},
		},
	})
}

func testAccRubrikAssignSLAConfig(vmName string, slaName string) string {
	return fmt.Sprintf(`
resource "rubrik_assign_sla" "assign_sla" {
	object_name = "%s"
	object_type = "vmware"
	sla_name    = "%s"
	timeout     = 15
	}
`, vmName, slaName)
}

func testAccSLAUnassigned() error {
	timeout := 15
	vmName := os.Getenv("RUBRIK_CDM_TEST_VM")
	rubrik := testAccProvider.Meta().(*rubrikcdm.Credentials)

	vmID, err := rubrik.ObjectID(vmName, "vmware", timeout)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	vmSummary, err := rubrik.Get("v1", fmt.Sprintf("/vmware/vm/%s", vmID), timeout)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	slaAssignment := vmSummary.(map[string]interface{})["slaAssignment"].(string)
	log.Printf("SLA: %s", slaAssignment)
	if slaAssignment != "Unassigned" {
		return fmt.Errorf("%s VM has an SLA Domain assigned: %s", vmName, slaAssignment)
	}

	return nil
}

func testAccCheckRubrikAssignSLADestroy(s *terraform.State) error {
	log.Printf("Running Post-check")
	err := testAccSLAUnassigned()
	if err != nil {
		return err
	}
	return nil
}

func testAccCheckRubrikAssignSLAPreCheck(t *testing.T) {
	log.Printf("Running Pre-check")
	err := testAccSLAUnassigned()
	if err != nil {
		t.Fatal(err.Error())
	}
}
