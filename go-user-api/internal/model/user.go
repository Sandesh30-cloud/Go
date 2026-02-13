package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"` //It controls JSON serialization.
	Age  int    `json:"age"`
}
