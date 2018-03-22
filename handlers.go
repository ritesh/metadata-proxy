package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!\n")
}

func iamInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!\n")
}
func iamSecurityCredentials(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if roleName, ok := vars["role-name"]; !ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusTeapot)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusTeapot, Text: "no role provided"}); err != nil {
			panic(err)
		}
		fmt.Printf("rolename is %v", roleName)
		return
	}
	//Fetch the creds somehow!
	creds := &IamCreds{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(creds); err != nil {
		panic(err)
	}
	return
}
