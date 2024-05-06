package structs

import (
	"time"
)

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
