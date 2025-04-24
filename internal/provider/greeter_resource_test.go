package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccGreeterGreeting_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { /* setup any necessary test preconditions */ },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
provider "greeter" {
  greeting_suffix = "!"
}

resource "greeter_greeting" "test" {
  name = "AccTest"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("greeter_greeting.test", "greeting", "Hello, AccTest!"),
				),
			},
			// Add an update test
			{
				Config: `
provider "greeter" {
  greeting_suffix = "?"
}

resource "greeter_greeting" "test" {
  name = "UpdatedTest"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("greeter_greeting.test", "greeting", "Hello, UpdatedTest?"),
				),
			},
		},
	})
}

// Define provider factories for testing
var providerFactories = map[string]func() (*schema.Provider, error){
	"greeter": func() (*schema.Provider, error) {
		return New(), nil
	},
}