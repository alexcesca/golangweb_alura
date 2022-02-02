package routes

import (
	"net/http"

	"github.com/loja/controllers"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
