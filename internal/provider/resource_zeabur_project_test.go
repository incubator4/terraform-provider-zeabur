package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccZeaburProjectResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + testAccZeaburProjectResourceConfig("test", "hkg1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("zeabur_project.test", "name", "test"),
					resource.TestCheckResourceAttr("zeabur_project.test", "region", "hkg1"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "zeabur_project.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"region", "last_updated"},
			},
			// Update and Read testing
			//{
			//	Config: testAccZeaburProjectResourceConfig("test", "hkg1"),
			//	Check: resource.ComposeAggregateTestCheckFunc(
			//		resource.TestCheckResourceAttr("zeabur_project.test", "name", "test"),
			//		resource.TestCheckResourceAttr("zeabur_project.test", "region", "hkg1"),
			//	),
			//},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccZeaburProjectResourceConfig(name, region string) string {
	return fmt.Sprintf(`
resource "zeabur_project" "test" {
  name   = %[1]q
  region = %[2]q
}
`, name, region)
}
