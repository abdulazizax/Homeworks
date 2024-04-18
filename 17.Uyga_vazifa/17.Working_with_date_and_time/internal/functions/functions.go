package functions

import (
	"fmt"
	store "homework/17.Working_with_date_and_time/storage"
)

func PrintEmployee(employee []store.Employee) {
	for _, v := range employee {
		fmt.Printf("Id: %v\nName: %v\nAge: %v\nPosition: %v\n\n", v.Id, v.Name, v.Age, v.Position)
	}
}
