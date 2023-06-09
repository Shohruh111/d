package controller

import (
	"app/models"
	"strings"

	"github.com/bxcodec/faker/v4"
)

func GenerateUser(count int) []models.User {
	var users []models.User
	for count >= 0 {
		users = append(users, models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count--
	}

	return users
}
func SearchPhoneNumber(users []models.User, number string) []models.User {
	var userNumber = []models.User{}
	for _, user := range users {
		if strings.HasPrefix(user.PhoneNumber, number) {
			userNumber = append(userNumber, user)
		}
	}
	return userNumber
}
