package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"

	_ "github.com/lib/pq"
)

var (
	drName     = "postgres"
	dbUrl      = "localhost"
	dbPort     = 5432
	dbName     = "sql"
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

type User struct {
	Id         uint
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Posts      []Post
}

type Post struct {
	Id         uint
	User_id    int
	Title      string
	Content    string
	Created_at time.Time
}

type Comment struct {
	Id         int
	Post_id    int
	User_id    int
	Comment    string
	Created_at time.Time
}

type LikePost struct {
	Id            int
	User_id       int
	Post_id       int
	Liked_user_id int
	Liked_at      time.Time
}

type LikeComment struct {
	Id            int
	User_id       int
	Post_id       int
	Comment_id    int
	Liked_user_id int
	Liked_at      time.Time
}

func main() {
	db := DBConnect()
	man := Manager{
		DB: db,
	}

	// 1. Create user
	// user := User{
	// 	Name: "Olim",
	// 	Email: "olim@gmail.com",
	// 	Password: "olim123",
	// }
	// new := man.CreateUser(&user)
	// pp.Println(new)

	// 2. Create post
	// post := Post{
	// 	User_id: 2,
	// 	Title: "post1",
	// 	Content: "Hello world!",
	// }
	// new := man.CreatePost(&post)
	// pp.Println(new)

	// 3. Create comment
	// comment := Comment{
	// 	User_id: 3,
	// 	Post_id: 2,
	// 	Comment: "Good job!",
	// }
	// new := man.CreateComment(&comment)
	// pp.Println(new)

	// 4. Get all Users
	new := man.GetAllUsers()
	for _, v := range new {
		fmt.Println(v.Id, ":")
		pp.Println(v)
	}

	// 5. Get Posts
	// new := man.GetPosts()
	// for _, v := range new {
	// 	pp.Println(v)
	// }

	// 6. Get User's post
	// new := man.GetAllUsers()
	// for _, v := range new {
	// 	man.GetUserPosts(&v)
	// 	fmt.Println(v.Id, ":")
	// 	for _, v1 := range v.Posts {
	// 		pp.Println(v1)
	// 	}
	// }

	// 7. Get Post's Comment
	// new := man.GetAllUsers()
	// for _, v := range new {
	// 	man.GetUserPosts(&v)
	// 	for _, v1 := range v.Posts {
	// 		post := man.GetPostComments(int(v1.Id))
	// 		fmt.Println(v1.Id, ":")
	// 		for _, v2 := range post {
	// 			pp.Println(v2)
	// 		}
	// 	}
	// }

	// 8. Creare Like for Post
	// new := man.CreateLikeToPost(1, 1, 3)
	// pp.Println(new)

	// 9. Create Like for Comment
	// new := man.CreateLikeToComment(1, 1, 1, 2)
	// pp.Println(*new)

}

// 1.
func (m *Manager) CreateUser(req *User) *User {

	user := User{}

	query := `
    INSERT INTO users(
      username,
      email,
	  password
    ) VALUES(
      $1,
      $2,
	  $3
    ) RETURNING
      id,
      username,
      email,
	  password,
	  created_at
  `
	err := m.DB.QueryRow(query, req.Name, req.Email, req.Password).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

// 2.
func (m *Manager) CreatePost(req *Post) *Post {
	var (
		post Post
	)

	query := `
    INSERT INTO posts(
      user_id,
      title,
      content
    )VALUES($1, $2, $3)
    RETURNING
      id,
      user_id,
      title,
      content,
      created_at
  `
	row := m.DB.QueryRow(query, req.User_id, req.Title, req.Content)

	err := row.Scan(&post.Id, &post.User_id, &post.Title, &post.Content, &post.Created_at)
	if err != nil {
		log.Fatal(err)
	}
	return &post
}

// 3.
func (m *Manager) CreateComment(req *Comment) Comment {
	var (
		com Comment
	)
	query := `
    INSERT INTO comments(
      user_id,
      post_id,
      content
    )VALUES($1, $2, $3)
    RETURNING
      id,
      user_id,
      post_id,
      content,
	  created_at
  `
	row := m.DB.QueryRow(query, req.User_id, req.Post_id, req.Comment)

	err := row.Scan(&com.Id, &com.User_id, &com.Post_id, &com.Comment, &com.Created_at)
	if err != nil {
		log.Fatal(err)
	}
	return com
}

// 4.
func (m *Manager) GetAllUsers() (users []User) {
	query := `
    SELECT * FROM users
  `
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		user := User{}

		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

// 5.
func (m *Manager) GetPosts() (posts []Post) {
	query := `
    SELECT 
      id, 
      user_id, 
      title, 
      content, 
      created_at
    FROM 
      posts
  `

	rows, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		p := Post{}
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	return posts
}

// 6.
func (m *Manager) GetUserPosts(user *User) {

	posts := []Post{}

	query := `
    SELECT 
      id, 
      user_id,
      title, 
      content, 
      created_at
    FROM 
      posts
    WHERE 
      user_id = $1
  `
	rows, err := m.DB.Query(query, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		p := Post{}
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	user.Posts = posts
}

// 7.
func (m *Manager) GetPostComments(post_id int) (comments []Comment) {

	query := `
    SELECT 
      id, 
      user_id,
      post_id, 
      content,
	  created_at
    FROM 
      comments
    WHERE 
      post_id = $1
  `

	rows, err := m.DB.Query(query, post_id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		c := Comment{}
		err := rows.Scan(&c.Id, &c.User_id, &c.Post_id, &c.Comment, &c.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, c)
	}
	return comments
}

// 8.
func (m *Manager) CreateLikeToPost(user_id, post_id, liked_user_id int) *LikePost {
	likes := LikePost{}

	query := `
    INSERT INTO likes_post(
      user_id,
      post_id,
	  liked_user_id
    ) VALUES(
      $1,
      $2,
	  $3
    ) RETURNING
      id,
      user_id,
      post_id,
	  liked_user_id,
	  liked_at
  `
	err := m.DB.QueryRow(query, user_id, post_id, liked_user_id).Scan(&likes.Id, &likes.User_id, &likes.Post_id, &likes.Liked_user_id, &likes.Liked_at)
	if err != nil {
		log.Fatal(err)
	}

	return &likes
}

// 9.
func (m *Manager) CreateLikeToComment(user_id, post_id, comment_id, liked_user_id int) *LikeComment {
	likes := LikeComment{}

	query := `
    INSERT INTO likes_comment(
      user_id,
      post_id,
	  comment_id,
	  liked_user_id
    ) VALUES(
      $1,
      $2,
	  $3,
	  $4
    ) RETURNING
      id,
      user_id,
      post_id,
	  comment_id,
	  liked_user_id,
	  liked_at
  `
	err := m.DB.QueryRow(query, user_id, post_id, comment_id, liked_user_id).Scan(&likes.Id, &likes.User_id, &likes.Post_id, &likes.Comment_id, &likes.Liked_user_id, &likes.Liked_at)
	if err != nil {
		log.Fatal(err)
	}

	return &likes
}