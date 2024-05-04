package function

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

type User struct {
	user_id    int
	username   string
	email      string
	password   string
	created_at time.Time
}

type Friendship struct {
	friendship_id int
	user_id       int
	friend_id     int
	status        string
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

func (m *Manager) GetAllUsers() (users []User) {
	query := `
    SELECT * FROM users;
  `
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		user := User{}

		err := row.Scan(&user.user_id, &user.username, &user.email, &user.password, &user.created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

func (m *Manager) GetAllFriendships() (friendships []Friendship) {
	query := `
    SELECT * FROM friendships;
  `
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		friendship := Friendship{}

		err := row.Scan(&friendship.friendship_id, &friendship.user_id, &friendship.friend_id, &friendship.status)
		if err != nil {
			log.Fatal(err)
		}
		friendships = append(friendships, friendship)
	}

	return friendships
}

func (m *Manager) AddToUsersTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("SQL/AddValuesToUsers.sql")
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

func (m *Manager) CreateUsersTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("SQL/CreateUsersTable.sql")
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

func (m *Manager) CreateFriendshipTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("SQL/CreateFriendshipsTable.sql")
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

func (m *Manager) AddToFriendshipsTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("SQL/AddValuesToFrienships.sql")
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

func (m *Manager) DropFriendshipsTableIfExists() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("SQL/DropFriendshipsTableIfExists.sql")
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

	query, err := m.ReadSqlFile("SQL/DropUsersTableIfExists.sql")
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
