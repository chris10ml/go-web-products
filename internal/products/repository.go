package products

import (
	"fmt"

	"github.com/chris10ml/go-web-products/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	LastID() (int, error)
	Store(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error)
	Update(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

var products []Product

// --- GET LAST ID ---

func (r *repository) LastID() (int, error) {
	var products []Product
	if err := r.db.Read(&products); err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

// --- API CRUD ---

func (r *repository) GetAll() ([]Product, error) {
	r.db.Read(&products)

	return products, nil
}

func (r *repository) Store(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error) {
	r.db.Read(&products)

	p := Product{id, name, color, price, stock, code, posted, dateCreated}
	products = append(products, p)

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) Update(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error) {
	u := Product{
		Name:        name,
		Color:       color,
		Price:       price,
		Stock:       stock,
		Code:        code,
		Posted:      posted,
		DateCreated: dateCreated,
	}

	updated := false

	r.db.Read(&products)

	for i := range products {
		if products[i].Id == id {
			u.Id = id
			products[i] = u
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return u, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var updatedProduct Product
	updated := false
	r.db.Read(&products)

	for i := range products {
        if products[i].Id == id {
            products[i].Name = name
            updated = true
            updatedProduct= products[i]
        }
    }

	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func (r *repository) Delete(id int) error {
	r.db.Read(&products)

	deleted := false
	for index, product := range products {
		if product.Id == id {
			products = append(products[:index], products[index+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Product %d not found", id)
	} else if err := r.db.Write(products); err != nil {
		return err
	}

	return nil
}
