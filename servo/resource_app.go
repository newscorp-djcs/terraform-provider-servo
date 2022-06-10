package servo

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

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
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// Elem: &schema.Resource{
				// 	Schema: map[string]*schema.Schema{
				// 		// "id": &schema.Schema{
				// 		// 	Type:     schema.TypeInt,
				// 		// 	Required: true,
				// 		// },
				// 		"handle": &schema.Schema{
				// 			Type:     schema.TypeString,
				// 			Required: true,
				// 		},
				// 		"source": &schema.Schema{
				// 			Type:     schema.TypeString,
				// 			Required: true,
				// 		},
				// 	},
				// },
				// Elem: &schema.Schema{
				// 	Type: schema.TypeString,
				// },
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

	// items := d.Get("app").([]interface{})
	ois := client.App{
		Handle: app_handle,
		Source: source,
	}

	appConfig := client.AppConfig{
		Region: region,
		Org:    org,
	}

	// for _, item := range items {
	// 	i := item.(map[string]interface{})

	// 	// co := i["handle"].([]interface{})[0]
	// 	// coffee := co.(map[string]interface{})

	// 	oi := App{
	// 		Handle: i["handle"].(string),
	// 		Source: i["source"].(string),
	// 	}

	// 	ois = append(ois, oi)
	// }

	o, err := c.CreateApp(ois, appConfig)
	if err != nil {
		os.WriteFile("logs", []byte(err.Error()), 0644)
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(o.ID))

	// resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appID := d.Id()
	appConfig := client.AppConfig{}
	app := client.App{}

	appRes, err := c.GetApp(appID, appConfig, app)
	if err != nil {
		return diag.FromErr(err)
	}

	appData := flattenAppAttributes(&appRes)
	if err := d.Set("appInfo", appData); err != nil {
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

func flattenAppAttributes(appRes client.AppsRes) []interface{} {
	a := make(map[string]interface{})
	a["id"] = appRes.ID
	a["context"] = appRes.Context
	a["handle"] = appRes.Handle
	a["updatedAt"] = appRes.UpdatedAt
	a["createdAt"] = appRes.CreatedAt
	a["source"] = appRes.Source

	return []interface{}{a}
}
