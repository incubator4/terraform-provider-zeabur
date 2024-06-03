package api

import (
	"github.com/zeabur/cli/pkg/api"
)

type Client struct {
	api.Client
}

func New(token string) *Client {
	client := api.New(token)
	return &Client{client}
}
