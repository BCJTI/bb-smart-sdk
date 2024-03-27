package bb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegisterBankSlip(t *testing.T) {
	slip := &BankSlip{
		NumeroConvenio:                       "",
		NumeroCarteira:                       0,
		NumeroVariacaoCarteira:               0,
		CodigoModalidade:                     0,
		DataEmissao:                          time.Time{},
		DataVencimento:                       time.Time{},
		ValorOriginal:                        0,
		ValorAbatimento:                      0,
		QuantidadeDiasProtesto:               0,
		IndicadorNumeroDiasLimiteRecebimento: "",
		NumeroDiasLimiteRecebimento:          0,
		CodigoAceite:                         "",
		CodigoTipoTitulo:                     "",
		DescricaoTipoTitulo:                  "",
		IndicadorPermissaoRecebimentoParcial: "",
		NumeroTituloBeneficiario:             "",
		CampoUtilizacaoBeneficiario:          "",
		NumeroTituloCliente:                  "",
		MensagemBloquetoOcorrencia:           "",
		Desconto:                             Discount{},
		JurosMora:                            Interest{},
		Multa:                                Fines{},
		Pagador:                              Teller{},
		BeneficiarioFinal:                    Recipient{},
		Email:                                "",
		IndicadorPix:                         "",
	}
	client.Authorize()
	register, err := client.Register(slip)
	assert.NoError(t, err)
	assert.NotEmpty(t, register.Beneficiario)
}
