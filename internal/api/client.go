package api

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

const (
	ZeaburServerURL = "https://gateway.zeabur.com"
	WebsocketURL    = "wss://gateway.zeabur.com"
)

const (
	ZeaburHTTPAPIEndpoint       = "https://gateway.zeabur.com/api/v1"
	ZeaburGraphQLAPIEndpoint    = ZeaburServerURL + "/graphql"
	WSSZeaburGraphQLAPIEndpoint = WebsocketURL + "/graphql"
)

type ObjectID string

type Client struct {
	*graphql.Client
	sub *graphql.SubscriptionClient
}

// NewGraphQLClientWithToken returns a new GraphQL client with the given token.
func NewGraphQLClientWithToken(token string) *graphql.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return graphql.NewClient(ZeaburGraphQLAPIEndpoint, httpClient)
}

func NewSubscriptionClient(token string) *graphql.SubscriptionClient {
	return graphql.NewSubscriptionClient(WSSZeaburGraphQLAPIEndpoint).
		WithConnectionParams(map[string]any{
			"authToken": token,
		})
}

func NewClient(token string) *Client {
	return &Client{
		NewGraphQLClientWithToken(token),
		NewSubscriptionClient(token),
	}
}
