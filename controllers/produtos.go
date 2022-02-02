package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Quanto faz um form value sempre retorna um String
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço:", err.Error())
		}
		quantidadeConvertida, errQtde := strconv.Atoi(quantidade)
		if errQtde != nil {
			log.Fatal("Erro na conversão da quantidade:", err.Error())
		}
		models.Novo(nome, descricao, precoConvertido, quantidadeConvertida)
		http.Redirect(w, r, "/", ST1031)
	}
}
