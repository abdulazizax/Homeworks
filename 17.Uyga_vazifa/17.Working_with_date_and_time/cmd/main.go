package main

import (
	"encoding/json"
	"fmt"
	funk "homework/17.Working_with_date_and_time/internal/functions"
	store "homework/17.Working_with_date_and_time/storage"
	"log"
	"os"
)

func main() {
	file, err := os.Open("17.Working_with_date_and_time/storage/employees.json")
	if err != nil {
		log.Fatal(err)
	}

	var employees []store.Employee

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&employees)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Json file dan ma'lumotlar o'qilib employees slice ga joylashtirildi: \n\n")
	funk.PrintEmployee(employees)

	// Indexi 3 ga teng bo'lgan employee ning yoshi 50 ga o'zgartirlidi
	for i, v := range employees {
		if v.Id == 3 {
			employees[i].Age = 50
			break
		}
	}

	fmt.Printf("Indexi 3 ga teng bo'lgan employee ning yoshi 50 ga o'zgartirlidi: \n\n")
	funk.PrintEmployee(employees)

	fmt.Println("Id si 6 ga teng bo'lgan employee ni ma'lumotlar bilan to'ldiring: ")
	var e store.Employee
	e.Id = 6
	fmt.Printf("Name: ")
	fmt.Scanln(&e.Name)
	fmt.Printf("Age: ")
	fmt.Scanln(&e.Age)
	fmt.Printf("Position: ")
	fmt.Scanln(&e.Position)
	employees = append(employees, e)

	fmt.Printf("\nemployees slice ga yangi employee haqida ma'lumotlar qo'shildi: \n\n")
	funk.PrintEmployee(employees)

	f, err := os.OpenFile("17.Working_with_date_and_time/storage/employees_new.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	err = encoder.Encode(employees)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("employees slice ichidagi ma'lumotlar `employees_new.json` file ga muvaffaqiyatli yozildi!!!")
}
