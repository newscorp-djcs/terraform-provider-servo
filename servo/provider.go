package servo

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// "golang.org/x/vuln/client"
	"local-servo-poc/client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// DefaultFunc: schema.EnvDefaultFunc("SERVO_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"servo_app": resourceApp(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"servo_app": dataSourceApps(),
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
