package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}

	u.ID = len(r.users) + 1

	r.users = append(r.users, &u)
	return &u, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email && u.Password == pass {
			return u, nil
		}
	}

	return nil, nil
}
