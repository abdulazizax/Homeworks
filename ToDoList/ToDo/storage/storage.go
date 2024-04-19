package storage

import "time"

type ToDo struct {
	Name string
	From time.Time
	To   time.Time
}

type User struct{
	User_Name string
	ToDoList []ToDo
}
