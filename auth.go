package bb

import (
	"net/http"
)

const oauthTokenPath = "/oauth/token"

type Client struct {
	ClientId     string
	ClientSecret string
	PublicKey    string
	AuthToken    AuthToken
	httpClient   *http.Client
}

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_id"`
	Scope       string `json:"scope"`
}

func NewClient(id, secret, token string) *Client {

	return &Client{
		ClientId:     id,
		ClientSecret: secret,
		AuthToken: AuthToken{
			AccessToken: token,
		},
		httpClient: http.DefaultClient,
	}

}

func (c *Client) Authorize() error {

	data := Params{
		"grant_type": "client_credentials",
		"scope":      "pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.lotes-info pagamentos-lote.devolvidos-info pagamentos-lote.cancelar-requisicao pagamentos-lote.transferencias-requisicao pagamentos-lote.transferencias-info pagamentos-lote.lotes-requisicao pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.lotes-requisicao pagamentos-lote.transferencias-info pagamentos-lote.transferencias-requisicao pagamentos-lote.cancelar-requisicao pagamentos-lote.devolvidos-info pagamentos-lote.lotes-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.pix-info pagamentos-lote.boletos-info",
	}

	return c.Post(oauthTokenPath, data, nil, &c.AuthToken)

}
