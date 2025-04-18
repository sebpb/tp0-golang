package main

import (
	"log"
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Servidor corriendo en el puerto 8080")

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
