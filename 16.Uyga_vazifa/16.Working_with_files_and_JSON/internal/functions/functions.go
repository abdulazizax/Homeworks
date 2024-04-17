package functions

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func Backup(firstFilePath, secondFilePath string) {
	firstFile, err := os.Open(firstFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	n := time.Now().Format("2006-01-02_150405")

	secondFilePath = secondFilePath+"file_backup_"+n+".txt"

	secondFile, err := os.Create(secondFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	io.Copy(secondFile, firstFile)

	NewVersion, err := os.OpenFile(firstFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	fmt.Printf("Filega yangi ma'lumot kiriting: ")
	var str string
	fmt.Scanln(&str)
	fmt.Fprintln(NewVersion, str)
}
