# Simple-go-service
A small service implemented via go to get list of books

## Usage:
    go run main.go handlers.go 
  
## Move to browser:
    http://localhost:8080/service/books
  
## Try via curl:
     curl -v -X GET http://localhost:8080/service/books
        Note: Unnecessary use of -X or --request, GET is already inferred.
        *   Trying ::1...
        * TCP_NODELAY set
        * Connected to localhost (::1) port 8080 (#0)
        > GET /service/books HTTP/1.1
        > Host: localhost:8080
        > User-Agent: curl/7.64.1
        > Accept: */*
        > 
        < HTTP/1.1 200 OK
        < Date: Wed, 25 Nov 2020 12:04:36 GMT
        < Content-Length: 102
        < Content-Type: text/plain; charset=utf-8
        < 
        Here is your list of books: 
                        1. Bunney and Rabbit
                        2. Lost in the Jungle
        * Connection #0 to host localhost left intact
                        3. Storms* Closing connection 0

## To run test Case:
    go test
