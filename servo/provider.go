package servo

import (
	"context"
	"os"

	"terraform-provider-servo/client"

	// "github.com/newscorp-djcs/terraform-provider-servo/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVO_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"servo_apps": dataSourceApps(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Token := d.Get("token").(string)
	Token := os.Getenv("SERVO_TOKEN")

	httpClient, _ := client.NewClient(nil, &Token)
	return httpClient, nil
}
