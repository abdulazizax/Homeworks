package functions

import (
	"fmt"
	"log"
	"time"

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

type Student struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	IsDeleted bool       `json:"is_deleted"`
}

type Course struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Price        float64   `json:"price"`
	StudentCount int       `json:"student_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	IsDeleted    bool      `json:"is_deleted"`
}

func (m *Manager) InsertIntoStudents(s Student) (Student, error) {
	var std Student

	tx, err := m.DB.Begin()
	if err != nil {
		return std, err
	}

	query := `
        INSERT INTO students (
            name,
            email)
        VALUES ($1, $2)
        RETURNING
			id,         
			name,       
			email,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
    `
	row := tx.QueryRow(query, s.Name, s.Email)

	err = row.Scan(&std.ID, &std.Name, &std.Email, &std.CreatedAt, &std.UpdatedAt, &std.DeletedAt, &std.IsDeleted)
	if err != nil {
		tx.Rollback()
		return std, err
	}

	err = tx.Commit()
	if err != nil {
		return std, err
	}

	return std, nil
}

func (m *Manager) InsertIntoCourse(c Course) (Course, error) {
	var crs Course

	tx, err := m.DB.Begin()
	if err != nil {
		return crs, err
	}

	query := `
        INSERT INTO courses (
            title,
            price)
        VALUES ($1, $2)
        RETURNING
			id,         
			title,       
			price,
			number_of_students,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
    `
	row := tx.QueryRow(query, c.Title, c.Price)

	err = row.Scan(&crs.ID, &crs.Title, &crs.Price, &crs.StudentCount, &crs.CreatedAt, &crs.UpdatedAt, &crs.DeletedAt, &crs.IsDeleted)
	if err != nil {
		tx.Rollback()
		return crs, err
	}

	err = tx.Commit()
	if err != nil {
		return crs, err
	}

	return crs, nil
}

func (m *Manager) AlterCourseTableById(c Course, id string) (Course, error) {
	var crs Course

	tx, err := m.DB.Begin()
	if err != nil {
		return crs, err
	}

	query := `
		UPDATE courses
		SET title = $1, price = $2, updated_at = $3
		WHERE id = $4
		RETURNING
			id,         
			title,       
			price,
			number_of_students,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	row := tx.QueryRow(query, c.Title, c.Price, time.Now(), id)

	err = row.Scan(&crs.ID, &crs.Title, &crs.Price, &crs.StudentCount, &crs.CreatedAt, &crs.UpdatedAt, &crs.DeletedAt, &crs.IsDeleted)
	if err != nil {
		tx.Rollback()
		return crs, err
	}

	err = tx.Commit()
	if err != nil {
		return crs, err
	}

	return crs, nil
}

func (m *Manager) AlterStudentTableById(s Student, id string) (Student, error) {
	var std Student

	tx, err := m.DB.Begin()
	if err != nil {
		return std, err
	}

	query := `
		UPDATE students
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
		RETURNING
			id,         
			name,       
			email,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	row := tx.QueryRow(query, s.Name, s.Email, time.Now(), id)

	err = row.Scan(&std.ID, &std.Name, &std.Email, &std.CreatedAt, &std.UpdatedAt, &std.DeletedAt, &std.IsDeleted)
	if err != nil {
		tx.Rollback()
		return std, err
	}

	err = tx.Commit()
	if err != nil {
		return std, err
	}

	return std, nil
}

func (m *Manager) GetAllStudents() (students []Student, err error) {
	query := `
	SELECT * FROM students
	`
	row, err := m.DB.Query(query)
	if err != nil {
		return students, err
	}

	for row.Next() {
		std := Student{}
		err = row.Scan(&std.ID, &std.Name, &std.Email, &std.CreatedAt, &std.UpdatedAt, &std.DeletedAt, &std.IsDeleted)
		if err != nil {
			return students, err
		}
		students = append(students, std)
	}
	return students, err
}

func (m *Manager) GetAllCourses() (courses []Course, err error) {
	query := `
	SELECT * FROM courses
	`
	row, err := m.DB.Query(query)
	if err != nil {
		return courses, err
	}

	for row.Next() {
		crs := Course{}
		err = row.Scan(&crs.ID, &crs.Title, &crs.Price, &crs.StudentCount, &crs.CreatedAt, &crs.UpdatedAt, &crs.DeletedAt, &crs.IsDeleted)
		if err != nil {
			return courses, err
		}
		courses = append(courses, crs)
	}
	return courses, err
}

func (m *Manager) DeleteFromStudentsById(id string) (Student, error) {
	var std Student

	tx, err := m.DB.Begin()
	if err != nil {
		return std, err
	}

	query := `
		UPDATE students
		SET deleted_at = $1, is_deleted = $2
		WHERE id = $3
		RETURNING
			id,         
			name,       
			email,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	row := tx.QueryRow(query, time.Now(), true, id)

	err = row.Scan(&std.ID, &std.Name, &std.Email, &std.CreatedAt, &std.UpdatedAt, &std.DeletedAt, &std.IsDeleted)
	if err != nil {
		tx.Rollback()
		return std, err
	}

	err = tx.Commit()
	if err != nil {
		return std, err
	}

	return std, nil
}

func (m *Manager) DeleteFromCourseById(id string) (Course, error) {
	var crs Course

	tx, err := m.DB.Begin()
	if err != nil {
		return crs, err
	}

	query := `
		UPDATE courses
		SET deleted_at = $1, is_deleted = $2
		WHERE id = $3
		RETURNING
			id,         
			title,       
			price,
			number_of_students,      
			created_at,
			updated_at,
			deleted_at,
			is_deleted
	`

	row := tx.QueryRow(query, time.Now(), true, id)

	err = row.Scan(&crs.ID, &crs.Title, &crs.Price, &crs.StudentCount, &crs.CreatedAt, &crs.UpdatedAt, &crs.DeletedAt, &crs.IsDeleted)
	if err != nil {
		tx.Rollback()
		return crs, err
	}

	err = tx.Commit()
	if err != nil {
		return crs, err
	}

	return crs, nil
}
