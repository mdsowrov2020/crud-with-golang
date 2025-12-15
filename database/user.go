package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (usr User) Store() User {
	if usr.ID != 0 {
		return usr
	}

	usr.ID = len(users) + 1

	users = append(users, usr)
	return usr
}

func Find(email, pass string) *User {
	for _, u := range users {
		if u.Email == email && u.Password == pass {
			return &u
		}
	}

	return nil
}
