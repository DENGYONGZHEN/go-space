package product

import (
	"database/sql"
	"go-api/types"
)

type ProductStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (s *ProductStore) GetProducts() ([]types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description, //这里的顺序要和type定义的struct的顺序一致，否则映射失败
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductStore) CreateProduct(p types.Product) error {

	_, err := s.db.Query("INSERT into products(name, description, image, price, quantity) Values (?,?,?,?,?)",
		p.Name, p.Description, p.Image, p.Price, p.Quantity)
	if err != nil {
		return err
	}
	return nil
}
