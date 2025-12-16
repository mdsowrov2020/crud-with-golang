package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"` // user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerB64 := base64UrlEncode(byteArrHeader)

	bytArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	paylodB64 := base64UrlEncode(bytArrData)

	message := headerB64 + "." + paylodB64
	bytArrSecret := []byte(secret)
	bytMessage := []byte(message)

	h := hmac.New(sha256.New, bytArrSecret)
	h.Write(bytMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := headerB64 + "." + paylodB64 + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
