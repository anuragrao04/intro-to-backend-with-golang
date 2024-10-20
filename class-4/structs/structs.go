package main

import "fmt"

type User struct {
	Username string
	Password string
}

func main() {
	user1 := User{
		Username: "user1",
		Password: "password1",
	}

	// _ := User{
	// 	Username: "user2",
	// 	Password: "password2",
	// }

	fmt.Println(user1.Username)
	fmt.Println(user1.Password)

	// users := []User{user1, user2}
	//
	//
	// for _, user := range users {
	// 	fmt.Println(user.Username)
	// }
}
