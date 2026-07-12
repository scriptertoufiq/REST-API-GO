package repo

import "errors"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		users: make([]User, 0),
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	for _, u := range r.users {
		if u.Email == user.Email {
			return nil, errors.New("email already exists")
		}
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	for i := range r.users {
		if r.users[i].Email == email && r.users[i].Password == pass {
			return &r.users[i], nil
		}
	}

	return nil, errors.New("user not found")
}
