package main

import "crud/cmd"

func main() {
	cmd.Serve()

	// secret := []byte("my-secret")
	// message := []byte("hello world")
	// h := hmac.New(sha256.New, secret)
	// h.Write(message)
	// text := h.Sum(nil)
	// fmt.Println(text)

	// jwt, err := util.CreateJWT("my-secret", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "Sowrov",
	// 	LastName:    "Khadem",
	// 	Email:       "test@mail.com",
	// 	IsShopOwner: true,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("JWT: ", jwt)
}
