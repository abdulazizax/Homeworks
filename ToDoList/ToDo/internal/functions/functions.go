package functions

import (
	"fmt"
	"os"
	"time"
	store "todo/ToDo/storage"
)

func Registration() {
	var name string
	fmt.Println("------------------------------------------------")
	fmt.Println("                   ToDoList")
	fmt.Println("------------------------------------------------")
	fmt.Printf("Ismingizni kiriting: ")
	fmt.Scanln(&name)
	var user store.User
	user.User_Name = name
	FirstWindow(user)
}

func FromToList(user store.User) {
	fmt.Printf("Boshlanish vaqtini kiriting (2024-04-17_07:08) : ")
	var from string
	fmt.Scanln(&from)
	check1, tm1 := FormatTime(from)
	if !check1 {
		fmt.Print("\nVaqtni noto'g'ri shaklda kiritdingiz, qaytadan urinib ko'ring!\n\n")
		FromToList(user)
	}

	fmt.Printf("Tugash vaqtini kiriting (2025-04-17_07:08) : ")
	var to string
	fmt.Scanln(&to)
	check2, tm2 := FormatTime(to)
	if !check2 {
		fmt.Print("\nVaqtni noto'g'ri shaklda kiritdingiz, qaytadan urinib ko'ring!\n\n")
		FromToList(user)
	}

	for _, val := range user.ToDoList {
		if val.From.After(tm1) && val.To.Before(tm2) {
			fmt.Printf("\nTopshiriq nomi: %v\nBoshlanish vaqti: %v\nTugash vaqti: %v\n\n\n", val.Name, val.From.Format("02 Jan 2006 15:04:05 MST"), val.To.Format("02 Jan 2006 15:04:05 MST"))
		}
	}
	ListOFAssignment(user)
}

func ListOFAssignment(user store.User) {
	fmt.Printf(`
[1] Barcha topshiriqlar ro'yxatini ko'rsatish
[2] Ma'lum bir vaqt oralig'idagi topshiriqlar ro'yxatini ko'rsatish
[3] Orqaga qaytish
[0] Exit

`)
	var choice int
	fmt.Printf("Tanlang: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		FullList(user)
	case 2:
		FromToList(user)
	case 3:
		FirstWindow(user)
	case 0:
		os.Exit(0)
	default:
		ListOFAssignment(user)
	}
}

func FullList(user store.User) {
	for _, val := range user.ToDoList {
		fmt.Printf("\nTopshiriq nomi: %v\nBoshlanish vaqti: %v\nTugash vaqti: %v\n\n\n", val.Name, val.From.Format("02 Jan 2006 15:04:05 MST"), val.To.Format("02 Jan 2006 15:04:05 MST"))
	}
	ListOFAssignment(user)
}

func FormatTime(t string) (bool, time.Time) {
	t = t + ":00"
	layout := "2006-01-02_15:04:05"

	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		fmt.Println("Xato yuz berdi:", err)
		return false, time.Now()
	}
	return true, parsedTime
}

func AddAssignment(user store.User) {
	var todo store.ToDo
	fmt.Printf("Topshiriqning nomini kiriting: ")
	fmt.Scanln(&todo.Name)

	fmt.Printf("Boshlanish vaqtini kiriting (2024-04-17_07:08) : ")
	var from string
	fmt.Scanln(&from)
	check1, tm1 := FormatTime(from)
	if !check1 {
		fmt.Print("\nVaqtni noto'g'ri shaklda kiritdingiz, qaytadan urinib ko'ring!\n\n")
		AddAssignment(user)
	}
	todo.From = tm1

	fmt.Printf("Tugash vaqtini kiriting (2025-04-17_07:08) : ")
	var to string
	fmt.Scanln(&to)
	check2, tm2 := FormatTime(to)
	if !check2 {
		fmt.Print("\nVaqtni noto'g'ri shaklda kiritdingiz, qaytadan urinib ko'ring!\n\n")
		AddAssignment(user)
	}
	todo.To = tm2

	user.ToDoList = append(user.ToDoList, todo)
	fmt.Printf(`
[1] Yangi topshiriq qo'shish
[2] Orqaga qaytish
[0] Exit

`)
	var choice int
	fmt.Printf("Tanlang: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		AddAssignment(user)
	case 2:
		FirstWindow(user)
	case 0:
		os.Exit(0)
	default:
		FirstWindow(user)
	}
}

func FirstWindow(user store.User) {
	fmt.Printf(`
[1] ToDoListga topshiriq qo'shish
[2] Topshiriqqlar ro'yxatini ko'rish
[0] Exit

`)
	var choice int
	fmt.Printf("Tanlang: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		AddAssignment(user)
	case 2:
		ListOFAssignment(user)
	case 0:
		os.Exit(0)
	default:
		FirstWindow(user)
	}
}
