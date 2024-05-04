package main

import (
	f "db/Function"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	db := f.DBConnect()

	man := f.Manager{
		DB: db,
	}
	man.DropFriendshipsTableIfExists()
	man.DropUsersTableIfExists()
	man.CreateUsersTable()
	man.CreateFriendshipTable()
	man.AddToUsersTable()
	man.AddToFriendshipsTable()

	users := man.GetAllUsers()
	fmt.Printf("\nAll users:\n\n")
	for _, v := range users {
		pp.Println(v)
	}

	friendships := man.GetAllFriendships()
	fmt.Printf("\nAll friendships:\n\n")
	for _, v := range friendships {
		pp.Println(v)
	}
}
