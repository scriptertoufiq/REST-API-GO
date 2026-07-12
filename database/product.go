package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

var productsList []Product

func Store(p Product) Product {
	p.ID = len(productsList) + 1
	productsList = append(productsList, p)
	return p
}

func List() []Product {
	return productsList
}

func Get(productID int) *Product {
	for _, product := range productsList {
		if product.ID == productID {
			return &product
		}
	}
	return nil

}

func Update(product Product) *Product {
	for i, p := range productsList {
		if p.ID == product.ID {
			productsList[i] = product
			return &productsList[i]
		}
	}
	return nil
}

func Delete(productID int) {
	var temptList []Product = make([]Product, 0)

	for _, product := range productsList {
		if product.ID != productID {
			temptList = append(temptList, product)
		}
	}
	productsList = temptList
}
