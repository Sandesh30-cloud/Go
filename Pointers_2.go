package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateAge(u *User) {
	u.Age = 30
}

func main() {
	user := User{Name: "Sandesh", Age: 22}
    updateAge(&user)
    fmt.Println(user.Age)
}



