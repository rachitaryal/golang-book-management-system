package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rachitaryal/golang-book-management-system/pkg/routes"
)

func main(){
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}