package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Requisição iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(time.Second * 5):
		// Log no servidor
		log.Println("Request processada com sucesso")
		// Log no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
