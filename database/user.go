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

func (user User) Store() User {
	if user.ID != 0 {
		return user
	}

	user.ID = len(users) + 1

	users = append(users, user)
	return user
}

func Find(email, pass string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == pass {
			return &user
		}
	}
	return nil
}
