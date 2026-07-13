package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"img_url" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	// Get(productID int) (*Product, error)
	List() ([]*Product, error)
	GetByID(productID int) *Product
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product
	err := r.dbCon.Select(&products, "SELECT id, title, description, price, img_url FROM products")
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			img_url
		)
		VALUES (
			:title,
			:description,
			:price,
			:img_url
		)
		RETURNING id;
	`
	var productID int
	row, err := r.dbCon.NamedQuery(query, p)
	if err != nil {
		return nil, err
	}

	if row.Next() {
		err = row.Scan(&productID)
	}
	p.ID = productID

	return &p, nil
}

// func (r *productRepo) Get(productID int) (*Product, error) {
// 	var product Product
// 	err := r.dbCon.Get(&product, "SELECT id, title, description, price, img_url FROM products WHERE id = $1", productID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

func (r *productRepo) GetByID(productID int) *Product {
	var product Product
	err := r.dbCon.Get(&product, "SELECT id, title, description, price, img_url FROM products WHERE id = $1", productID)
	if err != nil {
		return nil
	}
	return &product
}

func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
		UPDATE products
		SET title = :title,
			description = :description,
			price = :price,
			img_url = :img_url
		WHERE id = :id
	`
	result, err := r.dbCon.NamedExec(query, product)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no product found with ID %d", product.ID)
	}

	return &product, nil
}

func (r *productRepo) Delete(productID int) error {
	_, err := r.dbCon.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		return err
	}
	return nil
}
