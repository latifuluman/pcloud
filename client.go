package pcloud

import "net/http"

type PCloudClient struct {
	Auth   *string
	Client *http.Client
}

// NewClient create new PCloudClient
func NewClient() *PCloudClient {
	return &PCloudClient{
		Auth:   nil,
		Client: &http.Client{},
	}
}
