package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGreeting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreetingCreate,
		ReadContext:   resourceGreetingRead,
		UpdateContext: resourceGreetingUpdate,
		DeleteContext: resourceGreetingDelete,
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGreetingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*providerConfig)
	name := d.Get("name").(string)

	log.Printf("[INFO] Creating greeting for %s", name)

	d.SetId(name)
	greeting := fmt.Sprintf("Hello, %s%s", name, config.greetingSuffix)
	if err := d.Set("greeting", greeting); err != nil {
		return diag.FromErr(err)
	}

	return resourceGreetingRead(ctx, d, meta)
}

func resourceGreetingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*providerConfig)
	name := d.Id()

	log.Printf("[INFO] Reading greeting for %s", name)

	greeting := fmt.Sprintf("Hello, %s%s", name, config.greetingSuffix)
	if err := d.Set("greeting", greeting); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGreetingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*providerConfig)
	name := d.Get("name").(string)

	log.Printf("[INFO] Updating greeting for %s", name)

	// In a real provider, you might update the resource via an API
	greeting := fmt.Sprintf("Hello, %s%s", name, config.greetingSuffix)
	if err := d.Set("greeting", greeting); err != nil {
		return diag.FromErr(err)
	}

	return resourceGreetingRead(ctx, d, meta)
}

func resourceGreetingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Id()
	log.Printf("[INFO] Deleting greeting for %s", name)

	// In a real provider, you might delete the resource via an API
	d.SetId("")

	return nil
}