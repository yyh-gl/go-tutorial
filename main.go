package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    const Host = "localhost:5000" // Go的慣例：定数には型を指定しない

    router := mux.NewRouter()
    router.HandleFunc("/users/{id:[0-9]+}", usersHandler).Methods("GET")
    fmt.Println("Server Start >> " + Host)
    log.Fatal(http.ListenAndServe(Host, router))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User No.%v", vars["id"])
}
