package main

import (
	"net/http"
)

func BooksHandler(res http.ResponseWriter, req *http.Request){
        response := []byte(`Here is your list of books:
                            1. Bunney and Rabbit
                            2. Lost in the Jungle
                            3. Storms`)
        res.WriteHeader(200)
        res.Write(response)
}

