// Package service provides account service API
package service

import (
	"encoding/json"
	"net/http"

	"github.com/letrong/bm-account-service/dao"
	"github.com/letrong/bm-account-service/model"
)

// CreateAccountHandler create an account
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	dao.InsertNewAccount(mongoDb, logger, account)
	json.NewEncoder(w).Encode(account)
}

// GetAccountsHandler get all accounts
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{\"result\":\"OK\"}"))
}
