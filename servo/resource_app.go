package servo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// "strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Token string

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

type App struct {
	Metadata  Metadata `json:"metadata,omitempty"`
	Handle    string   `json:"handle,omitempty"`
	Context   string   `json:"context,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
	CreatedAt int64    `json:"created_at,omitempty"`
	Source    string   `json:"source,omitempty"`
}

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"app": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"handle": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	items := d.Get("app").([]interface{})
	ois := []App{}

	for _, item := range items {
		i := item.(map[string]interface{})

		// co := i["handle"].([]interface{})[0]
		// coffee := co.(map[string]interface{})

		oi := App{
			Handle: i["handle"].(string),
			Source: i["source"].(string),
		}

		ois = append(ois, oi)
	}

	Token := os.Getenv("SERVO_TOKEN")

	c.CreateApp(ois, Token)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// d.SetId(strconv.Itoa(o.ID))

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*Client)

	// // Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// orderID := d.Id()

	// order, err := c.GetOrder(orderID)
	// if err != nil {
	//   return diag.FromErr(err)
	// }

	// orderItems := flattenOrderItems(&order.Items)
	// if err := d.Set("items", orderItems); err != nil {
	//   return diag.FromErr(err)
	// }

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

func (c *Client) CreateApp(orderItems []App, Token string) (*App, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sorgs/dev/regions/virginia/apps", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, Token)
	if err != nil {
		return nil, err
	}

	app := App{}
	err = json.Unmarshal(body, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (c *Client) doRequest(req *http.Request, Token string) ([]byte, error) {
	token := Token

	// if authToken != nil {
	// 	token = *authToken
	// }

	req.Header.Set("token", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
