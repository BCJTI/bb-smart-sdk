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

const baseUrl = "https://oauth.hm.bb.com.br"

// Map data por post in the request
type Params url.Values

// Map extra request headers
type Headers map[string]string

func (c *Client) PostAuth(path string, params url.Values, model interface{}) error {
	contentType := "application/x-www-form-urlencoded"
	return c.execute("POST", path, contentType, strings.NewReader(params.Encode()), model)
}

/*
func (c *Client) Get(path string) error {
	contentType := "application/json; charset=utf-8"
	path += c.queryString()
	return c.execute(http.MethodGet, path, contentType, nil)
}

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
	/*return c.execute("POST", path, params, headers, model)*/
	return nil
}

func (c *Client) execute(method, path, contentType string, body io.Reader, model interface{}) error {

	// init vars
	var (
		url_api = baseUrl + path
		err     error
	)
	// validate method type
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return errors.New("rest: Not supported method")
	}
	request, err := http.NewRequest("POST", url_api, body)
	//request, err := http.NewRequest(http.MethodPost, url_api, body)
	if err != nil {
		return err
	}
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", contentType)

	// append access_token
	if path == oauthTokenPath {
		//request.Header.Add("content-type", "application/x-www-form-urlencoded")
		request.SetBasicAuth(c.ClientId, c.ClientSecret)

	}

	response, err := c.httpClient.Do(request)

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

	fmt.Print(content)

	// parse data
	return json.Unmarshal(content, model)

}

/*
func (c *Client) queryString() string {
	qs := c.params.Encode()
	if qs != "" {
		qs += "&"
	}
	return fmt.Sprintf("?%sgw-dev-app-key=%s",
		qs,
		c.credentials.AppKey,
	)
}
*/
