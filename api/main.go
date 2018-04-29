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

// HTTPError エラー
type HTTPError struct {
	Message string `json:"message"`
}

// GetAccount アカウント取得
func GetAccount(w http.ResponseWriter, r *http.Request) {
	account := Account{
		ID:   1,
		Name: "アカウント1",
	}
	respondJSON(w, http.StatusOK, account)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func main() {
	http.HandleFunc("/account", GetAccount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
