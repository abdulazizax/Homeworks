package main

import (
	"fmt"

	"github.com/k0kubun/pp"

	dbs "db/functions/db"
	mth "db/functions/methods"

	_ "github.com/lib/pq"
)

func main() {
	db := dbs.DBConnect()
	man := mth.Manager{
		DB: db,
	}

	// 1. Create user
	// user := str.User{
	// 	Name:     "Olim",
	// 	Email:    "olim@gmail.com",
	// 	Password: "olim123",
	// }
	// new := man.CreateUser(&user)
	// pp.Println(new)

	// 2. Create post
	// post := str.Post{
	// 	User_id: 2,
	// 	Title: "post1",
	// 	Content: "Hello world!",
	// }
	// new := man.CreatePost(&post)
	// pp.Println(new)

	// 3. Create comment
	// comment := str.Comment{
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
