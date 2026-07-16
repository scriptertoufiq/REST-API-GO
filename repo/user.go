package repo

import (
	"log"

	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
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

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email = $1 AND password = $2`
	var user domain.User
	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
