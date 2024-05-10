package functions

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	drName     = "postgres"
	dbUrl      = "localhost"
	dbPort     = 5432
	dbName     = "h28"
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

type Manager struct {
	DB *sqlx.DB
}

type Dataset struct {
	Id        int
	Generated int
}

func GetRandomNumber() int {
	return (rand.Intn(10) + 1) * 1000000
}

func (m *Manager) UpdateLarge_datasetTable(ctx context.Context) {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
	UPDATE large_dataset SET generated = 1 WHERE id = 1;
	`

	_, err = tx.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) InsertIntoLarge_dataset(ctx context.Context) {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
    INSERT INTO large_dataset(generated) VALUES($1);
    `
	_, err = tx.ExecContext(ctx, query, GetRandomNumber())
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) SelectFromLarge_dataset(ctx context.Context) (dataset []Dataset) {
	query := `
	SELECT 
  		id, 
  		generated
	FROM 
		large_dataset
	`

	row, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		dts := Dataset{}

		err := row.Scan(&dts.Id, &dts.Generated)
		if err != nil {
			log.Fatal(err)
		}
		dataset = append(dataset, dts)
	}

	return dataset
}

func (m *Manager) CreateLarge_datasetTable(ctx context.Context) {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}

	query := `
	CREATE TABLE large_dataset (
		id SERIAL PRIMARY KEY,
		generated INT
	);
	`
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) DropLarge_datasetTable(ctx context.Context) {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
	DROP TABLE IF EXISTS large_dataset;
	`

	_, err = tx.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}
