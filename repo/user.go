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
	Store(u User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	usersList []User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{
		usersList: []User{},
	}
	return repo
}

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
