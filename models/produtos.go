package models

import (
	"log"

	"github.com/loja/db"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
	Id         int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	todos, err := db.Query("select * from produtos ")
	if err != nil {
		log.Fatal(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for todos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = todos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.Id = id
		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func Novo(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	sqlInsert, errInc := db.Prepare("insert into produtos( nome, descricao,preco,quantidade) values ($1, $2, $3, $4)")
	if errInc != nil {
		log.Fatal("Erro no insert dos produtos: ", errInc.Error())
	}
	sqlInsert.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func Deletar(id string) {
	db := db.ConectaComBancoDeDados()
	sqlDelete, errInc := db.Prepare("delete from produtos where id = $1 ")
	if errInc != nil {
		log.Fatal("Erro no insert dos produtos: ", errInc.Error())
	}
	sqlDelete.Exec(id)
	defer db.Close()
}
