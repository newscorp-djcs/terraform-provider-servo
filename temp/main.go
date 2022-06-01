package main

import (
	"encoding/json"
	"fmt"
	"os"

	cError "github.com/coreos/etcd/error"
	"gopkg.in/resty.v1"
)

type AppsRes struct {
	Metadata  Metadata `json:"metadata,omitempty"`
	Handle    string   `json:"handle,omitempty"`
	Context   string   `json:"context,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
	CreatedAt int64    `json:"created_at,omitempty"`
	Source    string   `json:"source,omitempty"`
}
type Metadata struct {
	Stacks int `json:"stacks"`
}

type ArApps []AppsRes

const HostURL string = "https://next.onservo.com/api"

func main() {

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

	tempArrs := make([]AppsRes, 0)

	fmt.Print(resp.Body())

	ss := string(resp.Body())

	apps := json.Unmarshal([]byte(ss), &tempArrs)
	fmt.Printf("\n Apps: %v \n", apps)
}
