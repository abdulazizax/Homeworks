package main

import (
	"fmt"
	funk "homework/15.Interface/internal/functions"
)

func main() {
	filePath := "15.Interface/storage/sample.txt"
	info := funk.ReturnInfo(filePath)
 	result := funk.ReturnInfoAsSlice(info)

	fmt.Println("Slice ko'rinishi: ", result)

	fmt.Print("\n\nTartiblangan ko'rinishi:\n\n")
	for _, v := range result{
		fmt.Printf("Name: %s\nAge: %s\nOccupation: %s\n\n", v.Name, v.Age, v.Occupation)
	}
}
