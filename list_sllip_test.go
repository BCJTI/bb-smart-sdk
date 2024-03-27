package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSlip(t *testing.T) {
	client.Authorize()
	paramList := ListaBoletosParams{}
	list_slip, err := client.ListSlip(paramList)
	assert.NoError(t, err)
	assert.NotEmpty(t, list_slip.Boletos)
}
