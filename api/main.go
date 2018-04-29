package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Account アカウント
type Account struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HelloServer the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	account := Account{
		ID:   1,
		Name: "アカウント1",
	}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
