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
	todos, err := db.Query("select * from produtos order by id")
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}
func BuscaProduto(id int) Produto {
	db := db.ConectaComBancoDeDados()
	sqlConsulta, err := db.Query("select * from produtos where id = $1 ", id)
	if err != nil {
		log.Fatal("Erro no insert dos produtos: ", err.Error())
	}

	produtoBanco := Produto{}

	for sqlConsulta.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = sqlConsulta.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoBanco.Id = id
		produtoBanco.Nome = nome
		produtoBanco.Descricao = descricao
		produtoBanco.Preco = preco
		produtoBanco.Quantidade = quantidade
	}

	defer db.Close()
	return produtoBanco
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
func Update(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	sqlUpdate, errInc := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where  id = $5")
	if errInc != nil {
		log.Fatal("Erro no insert dos produtos: ", errInc.Error())
	}
	sqlUpdate.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
