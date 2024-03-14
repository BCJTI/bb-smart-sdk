package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const clientId = "eyJpZCI6IjExZGEzMDRhLTcyMzUtNDIyNC04MzMzLTQ5MmU0NDJkN2E2NyIsImNvZGlnb1B1YmxpY2Fkb3IiOjAsImNvZGlnb1NvZnR3YXJlIjo2OTk3Mywic2VxdWVuY2lhbEluc3RhbGFjYW8iOjF9"
const clientSecret = "eyJpZCI6ImYwMzk3YTAtMzRjYi00NGY2LWFlOGYtNWZiMjQ3ZWIyYzJlN2EiLCJjb2RpZ29QdWJsaWNhZG9yIjowLCJjb2RpZ29Tb2Z0d2FyZSI6Njk5NzMsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxLCJzZXF1ZW5jaWFsQ3JlZGVuY2lhbCI6MSwiYW1iaWVudGUiOiJob21vbG9nYWNhbyIsImlhdCI6MTY5MDk5NDQwODM4Mn0"
const token = ""

var client *Client

func init() {
	client = NewClient(clientId, clientSecret, token)
}

func TestAuthorize(t *testing.T) {
	err := client.Authorize()
	assert.NoError(t, err)
	assert.NotEmpty(t, client.AuthToken.AccessToken, "AccessToken cannot be empty")
	assert.NotEmpty(t, client.AuthToken.Scope, "Scope cannot be empty")
}
