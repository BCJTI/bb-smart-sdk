package bb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseUrlAuth = "https://oauth.hm.bb.com.br"
	baseUrl     = "https://api.hm.bb.com.br/cobrancas/v2"
)

// Map data por post in the request
type Params url.Values

// Map extra request headers
type Headers map[string]string

func (c *Client) PostAuth(path string, params url.Values, model interface{}) error {
	contentType := "application/x-www-form-urlencoded"
	return c.execute("POST", path, contentType, strings.NewReader(params.Encode()), model)
}

func (c *Client) Get(path string, params url.Values, model interface{}) error {
	contentType := "application/json; charset=utf-8"
	return c.execute(http.MethodGet, path, contentType, nil, model)
}

/*
func (c *Client) Post(path string, value interface{}) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	path += c.queryString()
	contentType := "application/json; charset=utf-8"
	return c.execute(http.MethodPost, path, contentType, bytes.NewReader(body))
}*/

// Execute POST requests
func (c *Client) Post(path string, params url.Values, model interface{}) error {
	contentType := "application/json; charset=utf-8"
	return c.execute("POST", path, contentType, strings.NewReader(c.setQueryParams(params)), model)
}

func (c *Client) setQueryParams(params url.Values) string {
	qs := params.Encode()
	if qs == "" {
		return fmt.Sprintf("?gw-dev-app-key=%s", c.AppKey)
	}

	return fmt.Sprintf("?%s&gw-dev-app-key=%s", qs, c.AppKey)
}

func (c *Client) execute(method, path, contentType string, body io.Reader, model interface{}) error {

	// init vars
	var (
		url_api = baseUrl + path + fmt.Sprintf("?gw-app-key=%s", c.AppKey)
		err     error
	)

	if path == oauthTokenPath {
		url_api = baseUrlAuth + path
	}

	// validate method type
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return errors.New("rest: Not supported method")
	}
	request, err := http.NewRequest(method, url_api, body)
	//request, err := http.NewRequest(http.MethodPost, url_api, body)
	if err != nil {
		return err
	}
	//request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", contentType)

	// append access_token or basic auth
	if path == oauthTokenPath {
		//request.Header.Add("content-type", "application/x-www-form-urlencoded")
		request.SetBasicAuth(c.ClientId, c.ClientSecret)
	} else {
		request.Header.Add("Authorization", "Bearer "+c.AuthToken.AccessToken)
		request.Header.Add("X-Application-Key", c.AppKey)
		request.Header.Add("X-Developer-Application-Key", c.AppKey)

	}

	response, err := c.httpClient.Do(request)
	fmt.Print(response)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// init MP custom error
	em := &ErrorMessage{}

	// check for error message
	if err = json.Unmarshal(content, em); err == nil && em.GetMessage() != "" {
		return em
	}

	// parse data
	return json.Unmarshal(content, model)
}
