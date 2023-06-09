package main

import (
	"app/controller"
	"fmt"
)

func main() {

	users := controller.GenerateUser(100)
	searchNumber := controller.SearchPhoneNumber(users, "102")
	for _, user := range searchNumber {
		fmt.Println(user)
	}

}
