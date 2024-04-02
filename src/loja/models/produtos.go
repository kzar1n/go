package models

import "loja/db"

type Produto struct {
	id              int
	nome, descricao string
	preco           float64
	quantidade      int
}

func (c *Produto) GetId() int {
	return c.id
}

func (c *Produto) GetNome() string {
	return c.nome
}

func (c *Produto) GetDescricao() string {
	return c.descricao
}

func (c *Produto) GetPreco() float64 {
	return c.preco
}

func (c *Produto) GetQuantidade() int {
	return c.quantidade
}

func (c *Produto) SetId(id int) {
	c.id = id
}

func (c *Produto) SetNome(nome string) {
	c.nome = nome
}

func (c *Produto) SetDescricao(descricao string) {
	c.descricao = descricao
}

func (c *Produto) SetPreco(preco float64) {
	c.preco = preco
}

func (c *Produto) SetQuantidade(quantidade int) {
	c.quantidade = quantidade
}

func Create(nome string, descricao string, preco float64, qtd int) {
	db := db.ConectaBanco()
	query, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(nome, descricao, preco, qtd)

	defer db.Close()
}

func List() []Produto {
	db := db.ConectaBanco()
	query, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.SetId(id)
		p.SetNome(nome)
		p.SetDescricao(descricao)
		p.SetPreco(preco)
		p.SetQuantidade(quantidade)

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func Remove(id int) {
	db := db.ConectaBanco()
	query, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)

	defer db.Close()
}

func GetProduto(id int) Produto {
	db := db.ConectaBanco()
	query, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.SetId(id)
		p.SetNome(nome)
		p.SetDescricao(descricao)
		p.SetPreco(preco)
		p.SetQuantidade(quantidade)
	}

	defer db.Close()
	return p
}

func Edit(id int) {
	db := db.ConectaBanco()
	query, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)

	defer db.Close()
}

func Update(id int, nome string, descricao string, preco float64, qtd int) {
	db := db.ConectaBanco()
	query, err := db.Prepare("update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id, nome, descricao, preco, qtd)

	defer db.Close()
}
