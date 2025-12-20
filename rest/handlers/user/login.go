package user

import (
	"encoding/json"
	"net/http"

	"crud/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	usr, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
	}

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internal server erro", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)
}
