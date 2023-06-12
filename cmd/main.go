package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"fmt"
)

func main() {
	cfg := config.Load()
	con := controller.NewController(&cfg)

	users := con.GenerateUser(100)
	con.Users = users

	var (
		dataLimit      int
		page           int
		searchName     string
		searchBirthday string
	)
	fmt.Println("Input Data Limit:")
	fmt.Scan(&dataLimit)

	paginationCount := len(con.Users) / dataLimit
	fmt.Println("Pages count: ", paginationCount)

	for {
		fmt.Println("Input page: ")
		fmt.Scan(&page)

		repUser := con.UserGetList(&models.UserGetListRequest{
			Offset: (page - 1) * dataLimit,
			Limit:  dataLimit,
		})

		for _, user := range repUser.Users {
			fmt.Println(user)
		}
		fmt.Println("Exit: 0")
		fmt.Scan(&page)
		if page == 0 {
			break
		}
	}
	for {
		fmt.Println("Enter a name for searching: ")
		fmt.Scan(&searchName)
		repUserName := con.GetListSearchByName(searchName)
		for _, user := range repUserName.Users {
			fmt.Println(user)
		}

		fmt.Println("Exit: 0")
		fmt.Scan(&page)
		if page == 0 {
			break
		}
	}

	for {
		fmt.Println("Enter a birthday for searching: ")
		fmt.Scan(&searchBirthday)

		repUserBirth := con.GetListSearchByBirthday(searchName)
		for _, user := range repUserBirth.Users {
			fmt.Println(user)
		}
	}

}
