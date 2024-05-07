package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

type Person struct {
	Id         int
	First_name string
	Lsst_name  string
}

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

func (m *Manager) ReadSqlFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var str strings.Builder

	for scanner.Scan() {
		str.WriteString(scanner.Text())
		str.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return str.String(), err
}

func (m *Manager) CreatePeopleTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("29.Working_with_network.Protocols/psql/create_people.sql")
	if err != nil {
		log.Fatal("Error reading SQL file: ", err)
	}

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) DropUsersTableIfExists() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("29.Working_with_network.Protocols/psql/drop_people_if_exits.sql")
	if err != nil {
		log.Fatal("Error reading SQL file: ", err)
	}

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) InserIntoPeopleTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("29.Working_with_network.Protocols/psql/insert_people.sql")
	if err != nil {
		log.Fatal("Error reading SQL file: ", err)
	}

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) CreateIndexOnFirstName() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
	CREATE INDEX idx_first_name ON people (first_name);
	`

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) DropIndexOnFirstName() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
	DROP INDEX IF EXISTS idx_first_name;
	`

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) CreateIndexOnFirstNameAndLastName() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query := `
	CREATE INDEX idx_first_name_last_name ON people (first_name, last_name);
	`

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
}

func (m *Manager) GetQueryPlanForAdrian() string {
	query := `
	EXPLAIN ANALYZE SELECT 
  		id, 
  		first_name, 
  		last_name
	FROM 
  		people 
	WHERE 
  		first_name = 'Adrian';
	`
	var queryPlan string

	row := m.DB.QueryRow(query)
	if err := row.Scan(&queryPlan); err != nil {
		log.Fatal("Error scanning: ", err)
	}

	return queryPlan
}

func (m *Manager) GetQueryPlanForAdrianGross() string {
	query := `
	EXPLAIN ANALYZE SELECT 
  		id, 
  		first_name, 
  		last_name
	FROM 
  		people 
	WHERE 
  		last_name = 'Gross' AND first_name = 'Adrian';
	`
	var queryPlan string

	row := m.DB.QueryRow(query)
	if err := row.Scan(&queryPlan); err != nil {
		log.Fatal("Error scanning: ", err)
	}

	return queryPlan
}

func (m *Manager) GetAdrianGross() (persons []Person) {
	query := `
	SELECT 
  		id, 
  		first_name, 
  		last_name
	FROM 
  		people 
	WHERE 
  		last_name = 'Gross' AND first_name = 'Adrian';
	`

	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		person := Person{}

		err := row.Scan(&person.Id, &person.First_name, &person.Lsst_name)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}

	return persons
}

func (m *Manager) GetAdrian() (persons []Person) {
	query := `
	SELECT 
  		id, 
  		first_name, 
  		last_name
	FROM 
  		people 
	WHERE 
  		first_name = 'Adrian';
	`

	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		person := Person{}

		err := row.Scan(&person.Id, &person.First_name, &person.Lsst_name)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}

	return persons
}
