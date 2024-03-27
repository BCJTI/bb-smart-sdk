package bb

import (
	"fmt"
	"strconv"
)

// CodeAccepted identifica se o boleto de cobrança foi aceito (reconhecimento da dívida pelo Pagador), sendo:
// "A" - ACEITE
// "N" - NAO ACEITE
type CodeAccepted string

const (
	Aceite    CodeAccepted = "A"
	NaoAceite CodeAccepted = "N"
)

// CodigoModalidade identifica a característica dos boletos dentro das modalidades de cobrança existentes no BB, sendo:
// 1 - SIMPLES
// 4 - VINCULADA
type CodeModality int

const (
	Simples   CodeModality = 1
	Vinculada CodeModality = 4
)

func (m CodeModality) String() string {
	if m == Simples || m == Vinculada {
		return strconv.Itoa(int(m))
	}
	return ""
}

// BoletoVencido indica se o boleto está vencido "S" ou não "N"
type ExpiredTicket string

const (
	Vencido    ExpiredTicket = "S"
	NaoVencido ExpiredTicket = "N"
)

// IndicadorPix informa se o boleto terá um QRCode Pix atrelado.
// Se informado caracter inválido,  assumirá 'N'.
// Domínios:
// 'S' - QRCODE DINAMICO;
// 'N' - SEM PIX;
// OUTRO - SEM PIX
type PixIndicator string

const (
	IndicadorPixSim PixIndicator = "S"
	IndicadorPixNao PixIndicator = "N"
)

// IndicadorSituacao indica a situação do boleto, sendo:
// A - boletos em ser
// B - boletos liquidados/baixados/protestados
type SituationIndicator string

const (
	BoletosEmSer                         SituationIndicator = "A"
	BoletosLiquidadosBaixadosProtestados SituationIndicator = "B"
)

// TipoDesconto indica cmo o desconto será concedido, sendo:
// 0 - SEM DESCONTO
// 1 - VLR FIXO ATE A DATA INFORMADA
// 2 - PERCENTUAL ATE A DATA INFORMADA
type TypeDescount int

const (
	SemDesconto TypeDescount = iota
	ValorFixo
	Percentual
)

// TipoJurosMora indica o código utilizado pela FEBRABAN para identificar o tipo de taxa de juros, sendo:
// 0 - DISPENSAR
// 1 - VALOR DIA ATRASO
// 2 - TAXA MENSAL
// 3 - ISENTO
type TypeInterest int

const (
	Dispensar TypeInterest = iota
	ValorDiaAtraso
	TaxaMensal
	Isento
)

// TipoMulta indica o código para identificação do tipo de multa para o Título de Cobrança, sendo:
// 0 - Sem multa
// 1 - Valor da Multa
// 2 - Percentual da Multa
type TypePenalty int

const (
	SemMulta TypePenalty = iota
	ValorDaMulta
	PercentualDaMulta
)

// TipoRegistro indica o tipo do registro do pagador, sendo:
// 1 - Pessoa Física
// 2 - Pessoa Jurídica
type TypeRecord int

const (
	PessoaFisica TypeRecord = iota + 1
	PessoaJuridica
)

// SituacaoBoleto indica o código da situação atual do boleto
type BillSituation int

const (
	SituacaoBoletoNormal                          BillSituation = iota + 1 // 01 NORMAL
	SituacaoBoletoMovimentoCartorio                                        // 02 MOVIMENTO CARTORIO
	SituacaoBoletoEmCartorio                                               // 03  EM CARTORIO
	SituacaoBoletoTituloComOcorrenciaCartorio                              // 04  TITULO COM OCORRENCIA DE CARTORIO
	SituacaoBoletoProtestadoEletronico                                     // 05  PROTESTADO ELETRONICO
	SituacaoBoletoLiquidado                                                // 06  LIQUIDADO
	SituacaoBoletoBaixado                                                  // 07  BAIXADO
	SituacaoBoletoTituloComPendenciaCartorio                               // 08  TITULO COM PENDENCIA DE CARTORIO
	SituacaoBoletoTituloProtestadoManual                                   // 09  TITULO PROTESTADO MANUAL
	SituacaoBoletoTituloBaixadoPagoEmCartorio                              // 10  TITULO BAIXADO/PAGO EM CARTORIO
	SituacaoBoletoTituloLiquidadoProtestado                                // 11  TITULO LIQUIDADO/PROTESTADO
	SituacaoBoletoTituloLiquidadoPagoEmCartorio                            // 12  TITULO LIQUID/PGCRTO
	SituacaoBoletoTituloProtestadoAguardandoBaixa                          // 13  TITULO PROTESTADO AGUARDANDO BAIXA
	SituacaoBoletoTituloEmLiquidacao                                       // 14  TITULO EM LIQUIDACAO
	SituacaoBoletoTituloAgendado                                           // 15  TITULO AGENDADO
	SituacaoBoletoTituloCreditado                                          // 16  TITULO CREDITADO
	SituacaoBoletoPagoEmCheque                                             // 17  PAGO EM CHEQUE - AGUARD.LIQUIDACAO
	SituacaoBoletoPagoParcialmente                                         // 18  PAGO PARCIALMENTE
	SituacaoBoletoPagoParcialmenteCreditado                                // 19  PAGO PARCIALMENTE CREDITADO
)

func (s BillSituation) String() string {
	if s > 0 {
		return fmt.Sprintf("%02d", s)
	}
	return ""
}
