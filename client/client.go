package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://next.onservo.com/api"

// var Token string = os.Getenv("SERVO_TOKEN")

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	// Auth       AuthStruct
}

type AppsRes struct {
	Metadata  Metadata `json:"metadata"`
	Handle    string   `json:"handle"`
	Context   string   `json:"context"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
	CreatedAt int64    `json:"created_at"`
	Source    string   `json:"source"`
	ID        int      `json:"id"`
}
type Metadata struct {
	Stacks int `json:"stacks"`
}

type App struct {
	// ID     int    `json:"id,omitempty"`
	Handle string `json:"handle,omitempty"`
	Source string `json:"source,omitempty"`
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// if token == nil {
	c.Token = *token
	// }

	// If username or password not provided, return empty client
	// if username == nil || password == nil {
	// 	return &c, nil
	// }

	// c.Auth = AuthStruct{
	// 	Username: *username,
	// 	Password: *password,
	// }

	// ar, err := c.SignIn()
	// if err != nil {
	// 	return nil, err
	// }

	// c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

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

func (c *Client) CreateApp(newApp App, Token string) (*AppsRes, error) {
	rb, err := json.Marshal(newApp)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sorgs/dev/regions/virginia/apps", "https://next.onservo.com/api/"), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("token", Token)

	// body, err := c.doRequest(req, Token)
	//---
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

	//---
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
