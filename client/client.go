package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HostURL - Default Servo URL
const HostURL string = "https://next.onservo.com/api/"

// var Token string = os.Getenv("SERVO_TOKEN")

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// if token == nil {
	c.Token = *token
	// }

	// c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
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

func (c *Client) CreateApp(newApp App, newAppConfig AppConfig) (*AppsRes, error) {
	rb, err := json.Marshal(newApp)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sorgs/%s/regions/%s/apps", HostURL, newAppConfig.Org, newAppConfig.Region), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("token", c.Token)

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	app := AppsRes{}
	err = json.Unmarshal(body, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (c *Client) GetApp(AppID string, appConfig AppConfig, app App) (AppsRes, error) {
	// req, err := http.NewRequest("GET", fmt.Sprintf("%sorgs/%s/regions/%s/apps/%s", HostURL, appConfig.Org, appConfig.Region, app.Handle), nil)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%sorgs/%s/regions/%s/apps/%s", HostURL, appConfig.Org, appConfig.Region, app.Handle), nil)
	// if err != nil {
	// 	return nil, err
	// }

	// body, err := c.doRequest(req)
	body, _ := c.doRequest(req)
	// if err != nil {
	// 	return nil, err
	// }

	appsRes := AppsRes{}
	// err = json.Unmarshal(body, &appsRes)
	json.Unmarshal(body, &appsRes)
	// if err != nil {
	// 	return nil, err
	// }

	return appsRes, nil
}
