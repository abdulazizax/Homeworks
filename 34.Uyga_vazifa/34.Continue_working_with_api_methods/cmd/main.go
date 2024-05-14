package main

import (
	funk "h34/34.Continue_working_with_api_methods/functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home/get_current_time", HomePage)
	log.Println("Server is listening on port :8080 : ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	str := funk.JSON_Time()
	log.Println(r.URL.Path)
	w.Write([]byte(str))
}
