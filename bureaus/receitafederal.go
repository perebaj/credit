// Package bureaus provides a client for the Receita Federal API.
package bureaus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const bureauName = "receita_federal"

// Empresa is a struct that represents a company in the Receita Federal.
type Empresa struct {
	NI                    string                `json:"ni"`
	TipoEstabelecimento   string                `json:"tipoEstabelecimento"`
	NomeEmpresarial       string                `json:"nomeEmpresarial"`
	NomeFantasia          string                `json:"nomeFantasia"`
	SituacaoCadastral     SituacaoCadastral     `json:"situacaoCadastral"`
	NaturezaJuridica      NaturezaJuridica      `json:"naturezaJuridica"`
	DataAbertura          string                `json:"dataAbertura"`
	CnaePrincipal         Cnae                  `json:"cnaePrincipal"`
	CnaeSecundarias       []Cnae                `json:"cnaeSecundarias"`
	Endereco              Endereco              `json:"endereco"`
	MunicipioJurisdicao   MunicipioCodigo       `json:"municipioJurisdicao"`
	Telefone              []Telefone            `json:"telefone"`
	CorreioEletronico     string                `json:"correioEletronico"`
	CapitalSocial         string                `json:"capitalSocial"`
	Porte                 string                `json:"porte"`
	SituacaoEspecial      string                `json:"situacaoEspecial"`
	DataSituacaoEspecial  string                `json:"dataSituacaoEspecial"`
	InformacoesAdicionais InformacoesAdicionais `json:"informacoesAdicionais"`
	ListaPeriodoSimples   []PeriodoSimples      `json:"listaPeriodoSimples"`
	Socios                []Socio               `json:"socios"`
}

// SituacaoCadastral is a struct that represents the situation of a company in the Receita Federal.
type SituacaoCadastral struct {
	Codigo string `json:"codigo"`
	Data   string `json:"data"`
	Motivo string `json:"motivo"`
}

// NaturezaJuridica is a struct that represents the legal nature of a company in the Receita Federal.
type NaturezaJuridica struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

// Cnae means Classificação Nacional de Atividades Econômicas.
type Cnae struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

// Endereco contains details about the address where the company is located.
type Endereco struct {
	TipoLogradouro string          `json:"tipoLogradouro"`
	Logradouro     string          `json:"logradouro"`
	Numero         string          `json:"numero"`
	Complemento    string          `json:"complemento"`
	CEP            string          `json:"cep"`
	Bairro         string          `json:"bairro"`
	Municipio      MunicipioCodigo `json:"municipio"`
	Pais           MunicipioCodigo `json:"pais"`
}

// MunicipioCodigo contains the code and description of a city that the company is located.
type MunicipioCodigo struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

// Telefone contains details about the phone number of the company.
type Telefone struct {
	DDD    string `json:"ddd"`
	Numero string `json:"numero"`
}

// InformacoesAdicionais is a struct that represents additional information about the company.
type InformacoesAdicionais struct {
	OptanteSimples string `json:"optanteSimples"`
	OptanteMei     string `json:"optanteMei"`
}

// PeriodoSimples is a struct that represents when the company opted for the Simples Nacional.
type PeriodoSimples struct {
	DataInicio string `json:"dataInicio"`
	DataFim    string `json:"dataFim"`
}

// Socio contains details about the partners of the company.
type Socio struct {
	TipoSocio          string             `json:"tipoSocio"`
	CPF                string             `json:"cpf"`
	Nome               string             `json:"nome"`
	Qualificacao       string             `json:"qualificacao"`
	DataInclusao       string             `json:"dataInclusao"`
	Pais               MunicipioCodigo    `json:"pais"`
	RepresentanteLegal RepresentanteLegal `json:"representanteLegal"`
}

// RepresentanteLegal contains details about the legal representative of the company.
type RepresentanteLegal struct {
	CPF          string `json:"cpf"`
	Nome         string `json:"nome"`
	Qualificacao string `json:"qualificacao"`
}

// RFClient is a client for the Receita Federal API.
type RFClient struct {
	Client *http.Client
	URL    string
	// OAuth2 access token
	Token string
}

// NewRFClient creates a new client for the Receita Federal API.
//
// This initialization contains the base URL for the Receita Federal API.
func NewRFClient(client *http.Client, token string) RFClient {
	return RFClient{
		Client: client,
		URL:    "https://apigateway.conectagov.estaleiro.serpro.gov.br",
		Token:  token,
	}
}

// Fetch retrieves information about a company from the Receita Federal.
// This fetch retrieves detailed information about a company based on its CNPJ.
func (c RFClient) Fetch(cnpj string, cpf string) (Empresa, error) {
	timer := time.Now()

	url := fmt.Sprintf("%s/api-cnpj-empresa/v2/empresa/%s", c.URL, cnpj)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Empresa{}, fmt.Errorf("could not create request to Receita Federal: %w. Url: %s", err, url)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Add("x-cpf-usuario", cpf)
	req.Header.Add("accept", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Empresa{}, fmt.Errorf("could not get data from Receita Federal: %w. Url: %s", err, url)
	}
	recordMetrics(resp.Status, timer, bureauName)

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return Empresa{}, fmt.Errorf("unexpected status code from Receita Federal: %d. Url: %s", resp.StatusCode, url)
	}

	var empresa Empresa
	if err := json.NewDecoder(resp.Body).Decode(&empresa); err != nil {
		return Empresa{}, fmt.Errorf("could not decode data from Receita Federal: %w", err)
	}
	return empresa, nil
}

// recordMetrics is a helper function to instrument requests that are made to the Receita Federal API.
func recordMetrics(statusCode string, timer time.Time, bureauName string) {
	duration := time.Since(timer).Seconds()
	// Record the duration of the request.
	bureauDuration.WithLabelValues(statusCode, bureauName).Observe(duration)
	// Record the request.
	bureauCounter.WithLabelValues(statusCode, bureauName).Inc()
}
