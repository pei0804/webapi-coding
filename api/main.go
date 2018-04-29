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

// HelloServer the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	account := Account{
		ID:   1,
		Name: "アカウント1",
	}
	res, err := json.Marshal(account)
	if err != nil {
		er := HTTPError{
			Message: err.Error(),
		}
		erres, _ := json.Marshal(er)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(erres)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
