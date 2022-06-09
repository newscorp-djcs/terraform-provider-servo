package servo

import (
	"context"
	"net/http"
	"os"

	// cError "github.com/coreos/etcd/error"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// "gopkg.in/resty.v1"
	"terraform-provider-servo/client"
)

var Token string

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// type App struct {
// 	// ID     int    `json:"id,omitempty"`
// 	Handle string `json:"handle,omitempty"`
// 	Source string `json:"source,omitempty"`
// }

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"app": &schema.Schema{
				Type:     schema.TypeMap,
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// Elem: &schema.Schema{
				// 	Type: schema.TypeString,
				// },
			},
			"org": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// Elem: &schema.Schema{
				// 	Type: schema.TypeString,
				// },
			},
		},
	}
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	app := d.Get("app").(map[string]interface{})
	region := d.Get("region").(string)
	org := d.Get("org").(string)

	app_handle := app["handle"].(string)
	source := app["source"].(string)

	// app_handle := d.Get("handle").(string)
	// source := d.Get("source").(string)

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

	// Token := os.Getenv("SERVO_TOKEN")
	// Token := c.Token

	// o, err := c.CreateApp(ois, Token)
	_, err := c.CreateApp(ois, appConfig)
	if err != nil {
		os.WriteFile("logs", []byte(err.Error()), 0644)
	}

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

// func (c *Client) CreateApp(newApp App, Token string) (*AppsRes, error) {
// 	rb, err := json.Marshal(newApp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%sorgs/dev/regions/virginia/apps", "https://next.onservo.com/api"), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("token", Token)

// 	// body, err := c.doRequest(req, Token)
// 	//---
// 	res, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
// 	}

// 	//---
// 	if err != nil {
// 		return nil, err
// 	}

// 	app := AppsRes{}
// 	err = json.Unmarshal(body, &app)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &app, nil
// }

// func (c *Client) doRequest(req *http.Request, Token string) ([]byte, error) {
// 	token := Token

// 	// if authToken != nil {
// 	// 	token = *authToken
// 	// }

// 	req.Header.Set("token", token)

// 	res, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
// 	}

// 	return body, err
// }
