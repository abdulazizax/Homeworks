package functions

import (
	"bufio"
	store "homework/15.Interface/storage"
	"log"
	"os"
	"strings"
)

func ReturnInfoAsSlice(info []map[string]string) []store.User {
	var result []store.User
	for _, val := range info {
		var user store.User
		for key, val1 := range val {
			switch key {
			case "Name":
				user.Name = val1
			case "Age":
				user.Age = val1
			case "Occupation":
				user.Occupation = val1
			}
		}
		result = append(result, user)
	}
	return result
}

func ReturnInfo(filePath string) []map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var usersInfo []map[string]string
	var currentUserInfo map[string]string

	currentUserInfo = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			result := strings.Split(line, ":")
			if len(result) == 2 {
				key := strings.TrimSpace(result[0])
				value := strings.TrimSpace(result[1])
				currentUserInfo[key] = value
			}
		}

		if len(currentUserInfo) == 3 {
			usersInfo = append(usersInfo, currentUserInfo)
			currentUserInfo = make(map[string]string)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return usersInfo
}
