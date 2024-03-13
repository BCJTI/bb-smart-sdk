#bb-smart-sdk

# BB SDK Boletos


SDK para o gateway de pagamento [BB](https://apoio.developers.bb.com.br/referency).

### URL's
## HH
 - https://oauth.sandbox.bb.com.br/oauth/token
 - https://oauth.sandbox.bb.com.br/oauth/token
 

## PRD
    - https://oauth.sandbox.bb.com.br/oauth/token
    - https://oauth.sandbox.bb.com.br/oauth/token
 

##### Endpoints implementados
- [ ] [Boleto - Lista](/boletos - GET)
- [ ] [Boleto - Registro](boletos/ - POST)
- [ ] [Boleto - Detalhe](/boletos/{id} - GET)
- [ ] [Boleto - Editar](/boletos/{id} - PATH)
- [ ] [Boleto - Baixa/Cancelamento](/boletos/{id}/baixar - POST)
- [ ] [Boleto - Cancelar Pix](/boletos/{id}/cancelar-pix - POST)
- [ ] [Boleto - Gerar Pix](/boletos/{id}/gerar-pix - POST)
- [ ] [Boleto - Consultar Pix](/boletos/{id}/pix - GET)

### Utilizando

```go
import (
	"github.com/BCJTI/bb-smart-sdk"
)

const clientId = ""
const clientSecret = ""
const publicKey = ""
const accessToken = "ACCESS_TOKEN"

func main() {
	client := mp.NewClient(clientId, clientSecret, publicKey, accessToken)

	pms, err := client.GetPaymentMethods()
	if err != nil {
		fmt.Println(err)
	}

	for _, pm := range pms {
		fmt.Println(pm.Name, pm.PaymentId)
	}
}
```




