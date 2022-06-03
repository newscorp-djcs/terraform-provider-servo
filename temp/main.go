package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	cError "github.com/coreos/etcd/error"
	"gopkg.in/resty.v1"
)

type AppsRes struct {
	Metadata  Metadata `json:"metadata,omitempty"`
	Handle    string   `json:"handle,omitempty"`
	Context   string   `json:"context,omitempty"`
	UpdatedAt int64    `json:"updatedAt,omitempty"`
	CreatedAt int64    `json:"createdAt,omitempty"`
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

	// // fmt.Print(resp.Body())

	fmt.Print(fmt.Sprintf("%s\n", resp.Body()))
	fmt.Print(reflect.TypeOf(resp.Body()), "\n")

	ss := string(resp.Body())

	// apps := json.Unmarshal([]byte(ss), &tempArrs)
	// fmt.Printf("\n Apps: %v \n", apps)

	fmt.Print(fmt.Sprintf("%s\n", resp.Body()))
	fmt.Print(reflect.TypeOf(resp.Body()), "\n")

	marshallErr := json.Unmarshal([]byte(ss), &tempArrs)
	if marshallErr != nil {
		panic(marshallErr)
	}
	fmt.Printf("\n Apps: %v \n", tempArrs)
}
