package servo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	cError "github.com/coreos/etcd/error"
	"gopkg.in/resty.v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-servo/client"
)

// type AppsRes struct {
// 	Metadata  Metadata `json:"metadata"`
// 	Handle    string   `json:"handle"`
// 	Context   string   `json:"context"`
// 	UpdatedAt int64    `json:"updated_at,omitempty"`
// 	CreatedAt int64    `json:"created_at"`
// 	Source    string   `json:"source"`
// 	ID        int      `json:"id"`
// }
// type Metadata struct {
// 	Stacks int `json:"stacks"`
// }

type ArApps []client.AppsRes

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
							Required: true,
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
							Required: true,
						},
					},
				},
			},
		},
	}
}

const HostURL string = "https://next.onservo.com/api"

func dataSourceAppsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	Token := os.Getenv("SERVO_TOKEN")

	client := resty.New().
		SetHostURL(HostURL).
		// SetTimeout(timeout).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			if r.IsSuccess() {
				return nil
			}

			return cError.NewError(r.StatusCode(), "error", 0)
		})

	// Create a request
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", Token).
		// SetBody(query).
		Get("/orgs/dev/regions/virginia/apps")

	if err != nil {
		fmt.Println(err)
	}

	tempArrs := ArApps{}
	apps := json.Unmarshal(resp.Body(), &tempArrs)
	fmt.Printf("\n Apps: %v \n", apps)

	if err := d.Set("apps", apps); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
