package main

import (
	"log"
	"net/http"

	"github.com/LucasWiman90/SimpleWeb/handlers"
)

func main() {
	http.HandleFunc("/view/", handlers.ViewHandler)
	http.HandleFunc("/edit/", handlers.EditHandler)
	http.HandleFunc("/save/", handlers.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
