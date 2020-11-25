package main

import (
	"net/http"
)

func main(){

	//Creating a mux
	mux := http.NewServeMux()


	//Mux is a router so registering API signature with handler
	mux.HandleFunc("/service/books", booksHandler)

	//Creating a server listening on :80
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	//Server to start listening
	server.ListenAndServe()

}

func booksHandler(res http.ResponseWriter, req *http.Request){
	response := []byte(`Here is your list of books: 
			    1. Bunney and Rabbit
			    2. Lost in the Jungle
			    3. Storms`)
	res.WriteHeader(200)
	res.Write(response)
}
