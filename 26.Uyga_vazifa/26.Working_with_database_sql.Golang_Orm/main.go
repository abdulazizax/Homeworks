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

type Products struct {
	product_name  string
	unit          string
	price         float64
	category_name string
	description   string
}

type Manager struct {
	DB *sqlx.DB
}

func GetAllProducts(m *Manager) (products []Products) {
	query := `
	SELECT 
    	product_name, 
    	unit, 
    	price, 
    	category_name, 
    	description 
	FROM 
    	(SELECT * FROM products
    	INNER JOIN categories 
    	ON products.category_id = categories.category_id) AS tbl
	WHERE 
    	category_name = 'Beverages';
	`
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal("Query error: ", err)
	}
	for row.Next() {
		prd := Products{}

		err := row.Scan(&prd.product_name, &prd.unit, &prd.price, &prd.category_name, &prd.description)
		if err != nil {
			log.Fatal("Scan error: ", err)
		}
		products = append(products, prd)
	}

	return products
}

func main() {
	m := Manager{
		DB: DBConnect(),
	}
	prd := GetAllProducts(&m)
	for _, v := range prd {
		fmt.Println(v)
	}
}
