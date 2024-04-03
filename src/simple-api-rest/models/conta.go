package models

type Conta struct {
	IdConta   string `json:"idConta"`
	Banco     string `json:"banco"`
	Agencia   string `json:"agencia"`
	Conta     string `json:"conta"`
	Dac       string `json:"dac"`
	TipoConta string `json:"tipoConta"`
}

func (c *Conta) Create(id string, banco string, agencia string, conta string, dac string, tipoConta string) *Conta {
	c.IdConta = id
	c.Banco = banco
	c.Agencia = agencia
	c.Conta = conta
	c.Dac = dac
	c.TipoConta = tipoConta
	return c
}
