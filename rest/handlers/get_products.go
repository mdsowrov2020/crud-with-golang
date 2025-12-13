// Package handlers

package handlers

import (
	"net/http"

	"crud/database"
	"crud/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
