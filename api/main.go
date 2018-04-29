package main

import (
	"encoding/json"
	"fmt"
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
	if r.Method != "GET" {
		er := HTTPError{
			Message: fmt.Sprintf("Not allowed method: %s", r.Method),
		}
		respondJSON(w, http.StatusMethodNotAllowed, er)
		return
	}
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
