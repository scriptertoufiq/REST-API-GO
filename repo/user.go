package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id;
	`
	var userId int
	row, err := r.dbCon.NamedQuery(query, user)
	if err != nil {
		log.Println("Error inserting user into the database:", err)
		return nil, err
	}

	if row.Next() {
		err = row.Scan(&userId)
	}
	user.ID = userId

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email = $1 AND password = $2`
	var user User
	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
