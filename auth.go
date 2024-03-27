package bb

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const oauthTokenPath = "/oauth/token"

type Client struct {
	ClientId     string
	ClientSecret string
	AppKey       string
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

	// Load client cert
	cert, err := tls.LoadX509KeyPair("server.crt", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile("ca-pfx.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return &Client{
		ClientId:     id,
		ClientSecret: secret,
		AppKey:       token,
		AuthToken: AuthToken{
			AccessToken: "",
		},
		httpClient: &http.Client{Transport: transport},
	}

}

func (c *Client) Authorize() error {

	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("scope", "cobrancas.boletos-requisicao cobrancas.boletos-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.lotes-info pagamentos-lote.devolvidos-info pagamentos-lote.cancelar-requisicao pagamentos-lote.transferencias-requisicao pagamentos-lote.transferencias-info pagamentos-lote.lotes-requisicao pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.lotes-requisicao pagamentos-lote.transferencias-info pagamentos-lote.transferencias-requisicao pagamentos-lote.cancelar-requisicao pagamentos-lote.devolvidos-info pagamentos-lote.lotes-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.pix-info pagamentos-lote.boletos-info")
	//params.Set("scope", "cobrancas.boletos-requisicao cobrancas.boletos-info")

	return c.PostAuth(oauthTokenPath, params, &c.AuthToken)

}
