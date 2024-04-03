package models

type Pagamento struct {
	Id          string  `json:"id"`
	Data        string  `json:"data"`
	Tipo        string  `json:"tipo"`
	Valor       float64 `json:"valor"`
	IdPagador   string
	Pagador     Pessoa `json:"pagador" gorm:"foreignKey:IdPessoa;references:IdPagador"`
	IdRecebedor string
	Recebedor   Pessoa `json:"recebedor" gorm:"foreignKey:IdPessoa;references:IdRecebedor"`
}

var Pagamentos []Pagamento

func (c *Pagamento) Create(id string, data string, tipo string, valor float64, pagador Pessoa, recebedor Pessoa) *Pagamento {
	c.Id = id
	c.Data = data
	c.Tipo = tipo
	c.Valor = valor
	c.Pagador = pagador
	c.Recebedor = recebedor
	return c
}
