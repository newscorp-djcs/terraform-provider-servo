package client

import "net/http"

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
	UpdatedAt int64    `json:"updatedAt,omitempty"`
	CreatedAt int64    `json:"createdAt"`
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
	// Region string `json:"region,omitempty"`
	// Org    string `json:"org,omitempty"`
}

type AppConfig struct {
	Region string `json:"region,omitempty"`
	Org    string `json:"org,omitempty"`
}
