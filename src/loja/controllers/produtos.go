package controllers

import (
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.ListaProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			panic(err.Error())
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			panic(err.Error())
		}

		models.CriaProdutos(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
