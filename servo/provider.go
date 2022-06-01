package servo

import (
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
	}
}
