package bb

import "net/http"

const oauthTokenPath = "/oauth/token"

type Client struct {
	ClientId       string
	ClientSecret   string
	ApplicationKey string
	AuthToken      AuthToken
	httpClient     *http.Client
}

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func NewClient(id, secret, key, scope string) *Client {

	return &Client{
		ClientId:       id,
		ClientSecret:   secret,
		ApplicationKey: key,
		httpClient:     http.DefaultClient,
	}

}

func (c *Client) Authorize() error {

	data := Params{
		"grant_type":    "client_credentials",
		"client_id":     c.ClientId,
		"client_secret": c.ClientSecret,
		"scope":         c.AuthToken.Scope,
	}

	return c.Post(oauthTokenPath, data, nil, &c.AuthToken)

}
