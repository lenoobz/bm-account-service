// Package service provides account service API
package service

import (
	"encoding/json"
	"net/http"

	"github.com/letrong/bm-account-service/dao"
	"github.com/letrong/bm-account-service/model"
)

// CreateAccountHandler creta a person
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	dao.InsertOneValue(mongoDb, account)
	json.NewEncoder(w).Encode(account)
}
