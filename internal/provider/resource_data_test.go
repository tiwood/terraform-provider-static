package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceData(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("static_data.testing", "data.some_data", "this is some data"),
					resource.TestCheckResourceAttr("static_data.testing", "data.more_data", "this is more data"),
					resource.TestCheckResourceAttr("static_data.testing", "output.some_data", "this is some data"),
					resource.TestCheckResourceAttr("static_data.testing", "output.more_data", "this is more data"),
				),
			},
			{
				Config: testAccResourceDataDataUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("static_data.testing", "data.some_data", "this is some data+changed"),
					resource.TestCheckResourceAttr("static_data.testing", "data.more_data", "this is more data+changed"),
					resource.TestCheckResourceAttr("static_data.testing", "output.some_data", "this is some data"),
					resource.TestCheckResourceAttr("static_data.testing", "output.more_data", "this is more data"),
				),
			},
			{
				Config: testAccResourceDataDataTriggersUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("static_data.testing", "triggers.test"),
					resource.TestCheckResourceAttr("static_data.testing", "data.some_data", "this is some data+changed"),
					resource.TestCheckResourceAttr("static_data.testing", "data.more_data", "this is more data+changed"),
					resource.TestCheckResourceAttr("static_data.testing", "output.some_data", "this is some data+changed"),
					resource.TestCheckResourceAttr("static_data.testing", "output.more_data", "this is more data+changed"),
				),
			},
		},
	})
}

const testAccResourceData = `
resource "static_data" "testing" {
  data = {
		some_data = "this is some data"
		more_data = "this is more data"
	}
}
`

const testAccResourceDataDataUpdate = `
resource "static_data" "testing" {
  data = {
		some_data = "this is some data+changed"
		more_data = "this is more data+changed"
	}
}
`

const testAccResourceDataDataTriggersUpdate = `
resource "static_data" "testing" {
  data = {
		some_data = "this is some data+changed"
		more_data = "this is more data+changed"
	}
	triggers = {
		test = "trigger1"
	}
}
`
