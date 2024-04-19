package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string, ch chan string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Faylni o'qishda xatolik: %s\n", filename)
		ch <- ""
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var str string
	for {
		line, err := reader.ReadString('\n')
		str += line
		if err != nil {
			break
		}
	}
	ch <- str
}

func main() {
	ch := make(chan string)
	filepath1 := "18.Parallelism_1st_part/Read_file/text1.txt"
	filepath2 := "18.Parallelism_1st_part/Read_file/text2.txt"
	filepath3 := "18.Parallelism_1st_part/Read_file/text3.txt"

	files := []string{filepath1, filepath2, filepath3}

	for _, filename := range files {
		go readFile(filename, ch)
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}
