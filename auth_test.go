package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const clientId = "eyJpZCI6IjExZGEzMDRhLTcyMzUtNDIyNC04MzMzLTQ5MmU0NDJkN2E2NyIsImNvZGlnb1B1YmxpY2Fkb3IiOjAsImNvZGlnb1NvZnR3YXJlIjo2OTk3Mywic2VxdWVuY2lhbEluc3RhbGFjYW8iOjF9"
const clientSecret = "eyJpZCI6ImYwMzk3YTAtMzRjYi00NGY2LWFlOGYtNWZiMjQ3ZWIyYzJlN2EiLCJjb2RpZ29QdWJsaWNhZG9yIjowLCJjb2RpZ29Tb2Z0d2FyZSI6Njk5NzMsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxLCJzZXF1ZW5jaWFsQ3JlZGVuY2lhbCI6MSwiYW1iaWVudGUiOiJob21vbG9nYWNhbyIsImlhdCI6MTY5MDk5NDQwODM4Mn0"
const scope = "pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.lotes-info pagamentos-lote.devolvidos-info pagamentos-lote.cancelar-requisicao pagamentos-lote.transferencias-requisicao pagamentos-lote.transferencias-info pagamentos-lote.lotes-requisicao pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.lotes-requisicao pagamentos-lote.transferencias-info pagamentos-lote.transferencias-requisicao pagamentos-lote.cancelar-requisicao pagamentos-lote.devolvidos-info pagamentos-lote.lotes-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-info pagamentos-lote.pagamentos-info pagamentos-lote.pagamentos-guias-sem-codigo-barras-requisicao pagamentos-lote.pagamentos-codigo-barras-info pagamentos-lote.boletos-requisicao pagamentos-lote.guias-codigo-barras-info pagamentos-lote.guias-codigo-barras-requisicao pagamentos-lote.transferencias-pix-info pagamentos-lote.transferencias-pix-requisicao pagamentos-lote.pix-info pagamentos-lote.boletos-info"
const applicationKey = "68b40d43431c569a0f833d7bf863e079"

var client *Client

func init() {
	client = NewClient(clientId, clientSecret, applicationKey, scope)
}

func TestAuthorize(t *testing.T) {
	err := client.Authorize()
	assert.NoError(t, err)
	assert.NotEmpty(t, client.AuthToken.AccessToken, "AccessToken cannot be empty")
	assert.NotEmpty(t, client.AuthToken.Scope, "RefreshToken cannot be empty")
}
