package main

import(
	funk "homework/16.Working_with_files_and_JSON/internal/functions"
)

func main() {
	firstFilePath := "16.Working_with_files_and_JSON/storage/firstFile.txt"
	secondFilePath := "16.Working_with_files_and_JSON/storage/"
	funk.Backup(firstFilePath, secondFilePath)
}