package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreeting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreetingRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name to greet",
			},
			"greeting": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The complete greeting message",
			},
		},
	}
}

func dataSourceGreetingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*providerConfig)
	name := d.Get("name").(string)

	log.Printf("[INFO] Reading greeting data source for %s", name)

	// Generate a unique ID for this data source
	d.SetId(fmt.Sprintf("greeting-%s", name))
	
	// Generate the greeting
	greeting := fmt.Sprintf("Hello, %s%s", name, config.greetingSuffix)
	if err := d.Set("greeting", greeting); err != nil {
		return diag.FromErr(err)
	}

	return nil
}