package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User No.%v", vars["id"])
}