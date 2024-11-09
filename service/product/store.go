package product

import (
	"database/sql"
	"fmt"
	"go-api/types"
	"strings"
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

func (s *ProductStore) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	//convert productIDs to []interface{}
	args := make([]interface{}, len(productIDs))

	for i, v := range productIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *ProductStore) UpdateProduct(p types.Product) error {

	_, err := s.db.Exec("UPDATE products SET name = ?,price = ?,image = ?,description = ?, quantity = ? WHERE id = ?",
		p.Name, p.Price, p.Image, p.Description, p.Quantity, p.ID)
	if err != nil {
		return err
	}
	return nil
}
