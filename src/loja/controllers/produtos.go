package controllers

import (
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.List())
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

		models.Create(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		panic(err.Error())
	}

	models.Remove(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		panic(err.Error())
	}

	temp.ExecuteTemplate(w, "Edit", append([]models.Produto{}, models.GetProduto(id)))
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))

		if err != nil {
			panic(err.Error())
		}

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

		models.Update(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
