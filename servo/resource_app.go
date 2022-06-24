package servo

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-servo/client"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				// Elem: &schema.Resource{
				Elem: &schema.Schema{
					Type: schema.TypeInt,
					// Schema: map[string]*schema.Schema{
					// 	"stacks": &schema.Schema{
					// 		Type:     schema.TypeInt,
					// 		Computed: true,
					// 	},
					// },
				},
			},
			"context": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"org": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	region := d.Get("region").(string)
	org := d.Get("org").(string)
	app_handle := d.Get("handle").(string)
	source := d.Get("source").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ac := client.App{
		Handle: app_handle,
		Source: source,
	}

	appConfig := client.AppConfig{
		Region: region,
		Org:    org,
	}

	ar, err := c.CreateApp(ac, appConfig)
	if err != nil {
		os.WriteFile("logs", []byte(err.Error()), 0644)
		return diag.FromErr(err)
	}

	appId := ar.Context + "/" + ar.Handle

	d.SetId(appId)

	resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appID := d.Id()
	// appConfig := client.AppConfig{}
	// app := client.App{}

	appResp, err := c.GetApp(appID)
	if err != nil {
		os.WriteFile("logsGETerr", []byte(err.Error()), 0644)
		return diag.FromErr(err)
	}

	metadata := make(map[string]int)
	// metadata := schema.NewSet()
	metadata["stacks"] = appResp.Metadata.Stacks

	// appData := flattenAppAttributes(appResp)
	if err := d.Set("metadata", metadata); err != nil {
		os.WriteFile("logs-set", []byte(err.Error()), 0644)
		return diag.FromErr(err)
	}
	if err := d.Set("handle", appResp.Handle); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("context", appResp.Context); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", appResp.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", appResp.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("source", appResp.Source); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceAppRead(ctx, d, m)
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

// func flattenAppAttributes(appRes *client.AppsRes) interface{} {
// 	a := make(map[string]interface{})
// 	// a["id"] = appRes.ID
// 	a["metadata"] = appRes.Metadata
// 	a["handle"] = appRes.Handle
// 	a["context"] = appRes.Context
// 	a["updatedAt"] = appRes.UpdatedAt
// 	a["createdAt"] = appRes.CreatedAt
// 	a["source"] = appRes.Source

// 	// return []interface{}{a}
// 	return a
// }
