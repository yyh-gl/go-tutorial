package main

import (
    "api-server-tutorial/handler"
    "fmt"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "net/http"
)

func main() {
    const Host = "localhost:5000"

    router := mux.NewRouter()
    router.HandleFunc("/users/{id:[0-9]+}", handler.UsersHandler).Methods("GET")

    fmt.Println("Server Start >> " + Host)
    log.Fatal(http.ListenAndServe(Host, router))
}
