package main

import (
	"net/http"
)

func main(){

	//Creating a mux
	mux := http.NewServeMux()


	//Mux is a router so registering API signature with handler
	mux.HandleFunc("/service/books", BooksHandler)

	//Creating a server listening on :80
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	//Server to start listening
	server.ListenAndServe()

}

