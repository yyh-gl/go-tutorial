package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "api-server-tutorial/handler"
)

func main() {
    const Host = "localhost:5000" // Go的慣例：定数には型を指定しない

    router := mux.NewRouter()
    router.HandleFunc("/users/{id:[0-9]+}", handler.UsersHandler).Methods("GET")
    fmt.Println("Server Start >> " + Host)
    log.Fatal(http.ListenAndServe(Host, router))
}
