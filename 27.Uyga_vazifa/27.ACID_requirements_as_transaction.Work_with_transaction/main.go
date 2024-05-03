package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	drName     = "postgres"
	dbUrl      = "localhost"
	dbPort     = 5432
	dbName     = "demo"
	dbUser     = "postgres"
	dbPassword = "abdulaziz1221"
)

func DBConnect() *sqlx.DB {
	dbStr := fmt.Sprintf(`
  host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		dbUrl, dbPort, dbUser, dbPassword, dbName)
	db, err := sqlx.Open(drName, dbStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Product struct {
	product_id   int
	product_name string
	category_id  int
	unit         string
	price        float64
}

func (m *Manager) CreateProduct(product_id int, product_name string, category_id int, unit string, price float64) Product {
	var prd Product

	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}

	query := `
        INSERT INTO products (
            product_id,
            product_name,
            category_id,
            unit,
            price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING
            product_id,
            product_name,
            category_id,
            unit,
            price;
    `
	row := tx.QueryRow(query, product_id, product_name, category_id, unit, price)

	err = row.Scan(&prd.product_id, &prd.product_name, &prd.category_id, &prd.unit, &prd.price)
	if err != nil {
		tx.Rollback()
		log.Fatal("Error in creating product: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	return prd
}

func (m *Manager) UpdateProductPrice(product_id int, price float64) Product {
	var prd Product

	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}

	query := `
        UPDATE products
        SET price = $1
        WHERE product_id = $2
        RETURNING product_id, product_name, category_id, unit, price;
    `
	row := tx.QueryRow(query, price, product_id)
	err = row.Scan(&prd.product_id, &prd.product_name, &prd.category_id, &prd.unit, &prd.price)
	if err != nil {
		tx.Rollback()
		log.Fatal("Error updating product price: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	return prd
}

func (m *Manager) DeleteProduct(product_id int) (Product, error) {
	var prd Product

	tx, err := m.DB.Begin()
	if err != nil {
		return prd, err
	}

	query := `
        DELETE FROM products
        WHERE product_id = $1
        RETURNING product_id, product_name, category_id, unit, price;
    `
	row := tx.QueryRow(query, product_id)
	err = row.Scan(&prd.product_id, &prd.product_name, &prd.category_id, &prd.unit, &prd.price)
	if err != nil {
		tx.Rollback()
		return prd, err
	}

	err = tx.Commit()
	if err != nil {
		return prd, err
	}

	return prd, nil
}

type Manager struct {
	DB *sqlx.DB
}

func main() {
	db := DBConnect()
	man := Manager{
		DB: db,
	}

	prd := Product{
		product_id:   100,
		product_name: "Qarshi suv",
		category_id:  1,
		unit:         "100 - 500 ml jars",
		price:        2.99,
	}

	newprd := man.CreateProduct(prd.product_id, prd.product_name, prd.category_id, prd.unit, prd.price)
	fmt.Println("Added product: ", newprd)
	newprd = man.UpdateProductPrice(prd.product_id, 3.99)
	fmt.Println("The price of the product will be changed to 3.99 so'm: ", newprd)
	newprd, err := man.DeleteProduct(prd.product_id)
	if err != nil {
		log.Fatal("Error deleting product: ", err)
	}
	fmt.Println("Deleted product: ", newprd)
}
