package main

import (
	"fmt"
	"log"
	"net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
	const dir = "files"
	const port = 8080

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))

	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", dir, port)


	handler := cors.Default().Handler(router)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), handler))
}
