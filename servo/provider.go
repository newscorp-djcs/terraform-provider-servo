package servo

import (
	"context"
	"os"
	// "terraform-provider-servo/client"
	"github.com/newscorp-djcs/terraform-provider-servo/client.git?ref=feature/addClient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"servo_apps": dataSourceApps(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	Token := os.Getenv("SERVO_TOKEN")

	httpClient := client.NewClient(nil, Token)
	return httpClient, nil
}
