package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"greeter_greeting": resourceGreeting(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"greeter_greeting": dataSourceGreeting(),
		},
		Schema: map[string]*schema.Schema{
			"greeting_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "!",
				Description: "The suffix to add to all greetings",
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfig contains any data needed from the provider configuration
type providerConfig struct {
	greetingSuffix string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := providerConfig{
		greetingSuffix: d.Get("greeting_suffix").(string),
	}

	return &config, nil
}