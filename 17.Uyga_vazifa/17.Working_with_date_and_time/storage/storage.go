package storage

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Position string `json:"position"`
}
