package methods

import (
	mn "db/functions/structs"
	"log"

	"github.com/jmoiron/sqlx"
)

type Manager struct {
	DB *sqlx.DB
}

// 1.
func (m *Manager) CreateUser(req *mn.User) *mn.User {

	user := mn.User{}

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
func (m *Manager) CreatePost(req *mn.Post) *mn.Post {
	var (
		post mn.Post
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
func (m *Manager) CreateComment(req *mn.Comment) mn.Comment {
	var (
		com mn.Comment
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
func (m *Manager) GetAllUsers() (users []mn.User) {
	query := `
    SELECT * FROM users
  `
	row, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		user := mn.User{}

		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

// 5.
func (m *Manager) GetPosts() (posts []mn.Post) {
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
		p := mn.Post{}
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	return posts
}

// 6.
func (m *Manager) GetUserPosts(user *mn.User) {

	posts := []mn.Post{}

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
		p := mn.Post{}
		err := rows.Scan(&p.Id, &p.User_id, &p.Title, &p.Content, &p.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	user.Posts = posts
}

// 7.
func (m *Manager) GetPostComments(post_id int) (comments []mn.Comment) {

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
		c := mn.Comment{}
		err := rows.Scan(&c.Id, &c.User_id, &c.Post_id, &c.Comment, &c.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, c)
	}
	return comments
}

// 8.
func (m *Manager) CreateLikeToPost(user_id, post_id, liked_user_id int) *mn.LikePost {
	likes := mn.LikePost{}

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
func (m *Manager) CreateLikeToComment(user_id, post_id, comment_id, liked_user_id int) *mn.LikeComment {
	likes := mn.LikeComment{}

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
