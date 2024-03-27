package bb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const bankSlipUrl = "/boletos"

func NumberCode(contrato, numero string) string {
	return fmt.Sprintf("000%s%010s", contrato, numero)
}

type BankSlip struct {
	NumeroConvenio                       string       `json:"numeroConvenio"`
	NumeroCarteira                       int          `json:"numeroCarteira"`
	NumeroVariacaoCarteira               int          `json:"numeroVariacaoCarteira"`
	CodigoModalidade                     CodeModality `json:"codigoModalidade"`
	DataEmissao                          time.Time    `json:"dataEmissao"`
	DataVencimento                       time.Time    `json:"dataVencimento"`
	ValorOriginal                        float64      `json:"valorOriginal"`
	ValorAbatimento                      float64      `json:"valorAbatimento"`
	QuantidadeDiasProtesto               int          `json:"quantidadeDiasProtesto"`
	IndicadorNumeroDiasLimiteRecebimento string       `json:"indicadorNumeroDiasLimiteRecebimento"`
	NumeroDiasLimiteRecebimento          int          `json:"numeroDiasLimiteRecebimento"`
	CodigoAceite                         CodeAccepted `json:"codigoAceite"`
	CodigoTipoTitulo                     string       `json:"codigoTipoTitulo"`
	DescricaoTipoTitulo                  string       `json:"descricaoTipoTitulo"`
	IndicadorPermissaoRecebimentoParcial string       `json:"indicadorPermissaoRecebimentoParcial"`
	NumeroTituloBeneficiario             string       `json:"numeroTituloBeneficiario"`
	CampoUtilizacaoBeneficiario          string       `json:"campoUtilizacaoBeneficiario"`
	NumeroTituloCliente                  string       `json:"numeroTituloCliente"`
	MensagemBloquetoOcorrencia           string       `json:"mensagemBloquetoOcorrencia"`
	Desconto                             Discount     `json:"desconto"`
	JurosMora                            Interest     `json:"jurosMora"`
	Multa                                Fines        `json:"multa"`
	Pagador                              Teller       `json:"pagador"`
	BeneficiarioFinal                    Recipient    `json:"beneficiarioFinal"`
	Email                                string       `json:"email"`
	IndicadorPix                         PixIndicator `json:"indicadorPix"`
}

func (b BankSlip) Validate() error {
	if len(strings.TrimSpace(b.NumeroConvenio)) != 7 {
		return errors.New("Número do convênio de cobrança inválido.")
	}
	if b.DataEmissao.Truncate(24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour)) {
		//Qualquer data, a partir da data atual no formato “dd.mm.aaaa”, entre aspas. CAMPO OBRIGATÓRIO.
		return errors.New("Data de emissão deve ser a partir da data atual.")
	}
	if b.DataVencimento.Truncate(24 * time.Hour).Before(b.DataEmissao.Truncate(24 * time.Hour)) {
		//Qualquer data >= dataEmissao, no formato “dd.mm.aaaa”, entre aspas. CAMPO OBRIGATÓRIO
		return errors.New("Data de vencimento deve ser maior, ou igual, a data de emissão.")
	}
	return nil
}

func (b BankSlip) MarshalJSON() ([]byte, error) {
	prefix := fmt.Sprintf("000%s", b.NumeroConvenio)
	if !strings.HasPrefix(b.NumeroTituloCliente, prefix) {
		b.NumeroTituloCliente = NumberCode(b.NumeroConvenio, b.NumeroTituloCliente)
	}

	type Alias BankSlip
	return json.Marshal(&struct {
		Alias
		DataEmissao    string `json:"dataEmissao"`
		DataVencimento string `json:"dataVencimento"`
	}{
		Alias:          (Alias)(b),
		DataEmissao:    b.DataEmissao.Format("02.01.2006"),
		DataVencimento: b.DataVencimento.Format("02.01.2006"),
	})
}

// RegistroBoleto é o retorno ao registar boleto
type RegisterBankSlip struct {
	Numero                 string `json:"numero"`
	NumeroCarteira         int    `json:"numeroCarteira"`
	NumeroVariacaoCarteira int    `json:"numeroVariacaoCarteira"`
	CodigoCliente          int64  `json:"codigoCliente"`
	LinhaDigitavel         string `json:"linhaDigitavel"`
	CodigoBarras           string `json:"codigoBarraNumerico"`
	NumeroContratoCobranca int64  `json:"numeroContratoCobranca"`
	Beneficiario           struct {
		Agencia              int64  `json:"agencia"`
		ContaCorrente        int64  `json:"contaCorrente"`
		TipoEndereco         int    `json:"tipoEndereco"`
		Logradouro           string `json:"logradouro"`
		Bairro               string `json:"bairro"`
		Cidade               string `json:"cidade"`
		CodigoCidade         int64  `json:"codigoCidade"`
		UF                   string `json:"uf"`
		CEP                  int64  `json:"cep"`
		IndicadorComprocavao string `json:"indicadorComprovacao"`
	} `json:"beneficiario"`
	QuantidadeOcorrenciasNegativacao string `json:"quantidadeOcorrenciasNegativacao"`
	//"listaOcorrenciasNegativacao": []
	QrCode QrCode `json:"qrCode"`
}

type QrCode struct {
	URL  string `json:"url"`
	TxId string `json:"txId"`
	Emv  string `json:"emv"`
}

// Desconto que será concedido no boleto
// Se tipo > 0, definir uma data de expiração do desconto, no formato "dd.mm.aaaa".
// Se tipo = 1, definir um valor de desconto >= 0.00 (formato decimal separado por ".").
// Se tipo = 2, definir uma porcentagem de desconto >= 0.00 (formato decimal separado por ".").
type Discount struct {
	Tipo          TypeDescount `json:"tipo"`
	DataExpiracao time.Time    `json:"dataExpiracao,omitempty"`
	Porcentagem   float64      `json:"porcentagem,omitempty"`
	Valor         float64      `json:"valor,omitempty"`
}

func (d Discount) MarshalJSON() ([]byte, error) {
	var data string
	if !d.DataExpiracao.IsZero() {
		data = d.DataExpiracao.Format("02.01.2006")
	}

	type Alias Discount
	return json.Marshal(&struct {
		Alias
		DataExpiracao string `json:"dataExpiracao,omitempty"`
	}{
		Alias:         (Alias)(d),
		DataExpiracao: data,
	})
}

// JurosMora define o valor de Juros que incide sobre o valor atual do boleto (valor do boleto - valor de abatimento)
// Se tipo = 1, definir um valor de desconto >= 0.00 (formato decimal separado por ".")
// Se tipo = 2, definir uma porcentagem de desconto >= 0.00 (formato decimal separado por ".")
type Interest struct {
	Tipo        TypeInterest `json:"tipo"`
	Porcentagem float64      `json:"porcentagem,omitempty"`
	Valor       float64      `json:"valor,omitempty"`
}

// Multa define o valor da Multa que incide sobre o valor atual do boleto (valor do boleto - valor de abatimento).
// Se tipo = 0 (zero) os campos “DATA DE MULTA”, “PERCENTUAL DE MULTA” e “VALOR DA MULTA” não devem ser informados ou ser informados iguais a ‘0’ (zero).
type Fines struct {
	Tipo        TypePenalty `json:"tipo"`
	Data        time.Time   `json:"data,omitempty"`
	Porcentagem float64     `json:"porcentagem,omitempty"`
	Valor       float64     `json:"valor,omitempty"`
}

func (f Fines) MarshalJSON() ([]byte, error) {
	var data string
	if !f.Data.IsZero() {
		data = f.Data.Format("02.01.2006")
	}

	type Alias Fines
	return json.Marshal(&struct {
		Alias
		Data string `json:"data,omitempty"`
	}{
		Alias: (Alias)(f),
		Data:  data,
	})
}

type Recipient struct {
	Nome            string     `json:"nome"`
	NumeroInscricao string     `json:"numeroInscricao"`
	TipoInscricao   TypeRecord `json:"tipoInscricao"`
}

func (b Recipient) MarshalJSON() ([]byte, error) {
	type Alias Recipient
	return json.Marshal(&struct {
		Alias
		Nome string `json:"nome"`
	}{
		Alias: (Alias)(b),
		Nome:  substring(b.Nome, 30),
	})
}

type Teller struct {
	Nome            string     `json:"nome"`
	NumeroInscricao string     `json:"numeroInscricao"`
	TipoInscricao   TypeRecord `json:"tipoInscricao"`
	Telefone        string     `json:"telefone"`
	Endereco        string     `json:"endereco"`
	Bairro          string     `json:"bairro"`
	CEP             string     `json:"cep"`
	Cidade          string     `json:"cidade"`
	UF              string     `json:"uf"`
}

func (p Teller) MarshalJSON() ([]byte, error) {
	type Alias Teller
	return json.Marshal(&struct {
		Alias
		Nome     string `json:"nome"`
		Endereco string `json:"endereco"`
		Bairro   string `json:"bairro"`
		Cidade   string `json:"cidade"`
	}{
		Alias:    (Alias)(p),
		Nome:     substring(p.Nome, 30),
		Endereco: substring(p.Endereco, 30),
		Bairro:   substring(p.Bairro, 30),
		Cidade:   substring(p.Cidade, 30),
	})
}

func substring(s string, l int) string {
	if len(s) > l {
		return s[0:l]
	}
	return s
}

func (c *Client) Register(bankSlip *BankSlip) (*RegisterBankSlip, error) {

	if bankSlip.IndicadorPermissaoRecebimentoParcial == "" {
		bankSlip.IndicadorPermissaoRecebimentoParcial = "N"
	}

	err := c.Post(bankSlipUrl, url.Values{}, bankSlip)

	register := &RegisterBankSlip{}

	return register, err

}
