package servo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"
	"strconv"
	"time"

	// cError "github.com/coreos/etcd/error"
	// "gopkg.in/resty.v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AppsRes struct {
	Metadata  Metadata `json:"metadata"`
	Handle    string   `json:"handle"`
	Context   string   `json:"context"`
	UpdatedAt int64    `json:"updatedAt,omitempty"`
	CreatedAt int64    `json:"createdAt"`
	Source    string   `json:"source"`
}
type Metadata struct {
	Stacks int `json:"stacks"`
}

type ArApps []AppsRes

func dataSourceApps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppsRead,
		Schema: map[string]*schema.Schema{
			"apps": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"handle": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"context": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						// "updated_at": &schema.Schema{
						// 	Type:     schema.TypeString,
						// 	Computed: true,
						// },
						// "created_at": &schema.Schema{
						// 	Type:     schema.TypeInt,
						// 	Computed: true,
						// },
						"source": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

const HostURL string = "https://next.onservo.com/api"

func dataSourceAppsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// org := d.Get("org").(string)
	// region := d.Get("region").(string)
	// app_handle := d.Get("app_handle").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// req, err := http.NewRequest("GET", fmt.Sprintf("%sorgs/dev/regions/${region}/apps/${app_handle}", "https://next.onservo.com/api/"), nil)
	req, err := http.NewRequest("GET", fmt.Sprintf("%sorgs/${org}/regions/virginia/apps/admin-djcss", "https://next.onservo.com/api/"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	apps := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&apps)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("apps", apps); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
