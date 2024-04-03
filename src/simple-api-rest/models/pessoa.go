package models

type Pessoa struct {
	IdPessoa      string `json:"idPessoa"`
	Nome          string `json:"nome"`
	Documento     string `json:"documento"`
	TipoDocumento string `json:"tipoDocumento"`
	IdConta       string `json:"-"`
	Conta         Conta  `json:"conta" gorm:"foreignKey:IdConta;references:IdConta"`
}

func (c *Pessoa) Create(nome string, documento string, tipoDocumento string, conta Conta) *Pessoa {
	c.Nome = nome
	c.Documento = documento
	c.TipoDocumento = tipoDocumento
	c.Conta = conta
	return c
}
