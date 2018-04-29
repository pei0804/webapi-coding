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

// ListAccount アカウント取得
func ListAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		er := HTTPError{
			Message: fmt.Sprintf("Not allowed method: %s", r.Method),
		}
		respondJSON(w, http.StatusMethodNotAllowed, er)
		return
	}
	accounts := []Account{
		Account{
			ID:   1,
			Name: "アカウント1",
		},
		Account{
			ID:   2,
			Name: "アカウント2",
		},
	}
	respondJSON(w, http.StatusOK, accounts)
}

// ShowAccount アカウント取得
func ShowAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		er := HTTPError{
			Message: fmt.Sprintf("Not allowed method: %s", r.Method),
		}
		respondJSON(w, http.StatusMethodNotAllowed, er)
		return
	}
	var id int
	if _, err := fmt.Sscanf(r.URL.Path, "/accounts/%d", &id); err != nil {
		er := HTTPError{
			Message: fmt.Sprintf("Invalid id: %s", err),
		}
		respondJSON(w, http.StatusMethodNotAllowed, er)
		return
	}
	account := Account{
		ID:   id,
		Name: fmt.Sprint("アカウント", id),
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
	// accounts/id 単一ユーザーを取得する
	http.HandleFunc("/accounts/", ShowAccount)
	// accounts リスト取得
	http.HandleFunc("/accounts", ListAccount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
