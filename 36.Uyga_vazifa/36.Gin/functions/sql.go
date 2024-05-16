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
	dbName     = "h35"
	dbUser     = "postgres"
	dbPassword = "abdulaziz1221"
)

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
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

func (m *Manager) CreateArtistTable() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	defer tx.Rollback()

	query, err := m.ReadSqlFile("36.Gin/psql/CreateAlbumTable.sql")
	if err != nil {
		log.Fatal("Error reading sql file: ", err)
	}

	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal("Error executing sql query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error commiting transaction: ", err)
	}
}

func (m *Manager) SelectAlbumsById(id int) (Album, error) {
	query := `
	SELECT * FROM albums WHERE id = $1
	`
	row := m.DB.QueryRow(query, id)

	album := Album{}

	err := row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		return album, err
	}

	return album, nil
}

func (m *Manager) SelectAlbums() (albums []Album) {
	query := `
	SELECT * FROM albums
	`
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal("Error executing sql query: ", err)
	}

	for row.Next() {
		album := Album{}
		err := row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			log.Fatal("Error scanning sql query: ", err)
		}
		albums = append(albums, album)
	}
	return albums
}

func (m *Manager) AlterColumnById(alb Album, id int) (Album, error) {
	result := Album{}

	tx, err := m.DB.Begin()
	if err != nil {
		fmt.Println("1 er")
		return result, err
	}

	query := `
		UPDATE albums
		SET title = $1, artist = $2, price = $3
		WHERE id = $4
		RETURNING
			id,         
			title,       
			artist,      
			price
	`

	row := tx.QueryRow(query, alb.Title, alb.Artist, alb.Price, id)

	err = row.Scan(&result.Id, &result.Title, &result.Artist, &result.Price)
	if err != nil {
		tx.Rollback()
		fmt.Println("2 er")
		return result, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("3 er")
		return result, err
	}

	return result, nil
}

func (m *Manager) InsertIntoAlbums(alb Album) Album {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}

	query := `
		INSERT INTO albums (
			title, 
			artist, 
			price) 
		VALUES(	$1, $2, $3)
		RETURNING
			id,         
			title,       
			artist,      
			price
	`

	row := tx.QueryRow(query, alb.Title, alb.Artist, alb.Price)
	album := Album{}

	err = row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		tx.Rollback()
		log.Fatal("Error scanning sql query: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	return album
}
