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
