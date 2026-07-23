package repo

import (
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product interface {
	product.ProductRepo
}

type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) *productRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.dbCon.Select(&products, "SELECT id, title, description, price, img_url FROM products LIMIT $1 OFFSET $2", limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
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
	rows, err := r.dbCon.NamedQuery(query, p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&productID); err != nil {
			return nil, err
		}
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

func (r *productRepo) GetByID(productID int) *domain.Product {
	var product domain.Product
	err := r.dbCon.Get(&product, "SELECT id, title, description, price, img_url FROM products WHERE id = $1", productID)
	if err != nil {
		return nil
	}
	return &product
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
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

func (r *productRepo) Count() (int64, error) {
	var count int64
	err := r.dbCon.Get(&count, "SELECT COUNT(*) FROM products")
	if err != nil {
		return 0, err
	}
	return count, nil
}
