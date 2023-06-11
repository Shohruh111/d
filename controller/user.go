package controller

import (
	"app/models"
	"log"
	"strings"

	"github.com/bxcodec/faker/v4"
)

func (c *Controller) GenerateUser(count int) []*models.User {
	var users []*models.User
	for count >= 0 {
		users = append(users, &models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
			Birthday:    faker.Date(),
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

func (c *Controller) UserGetList(req *models.UserGetListRequest) *models.UserGetListResponse {
	log.Printf("UserGetList req: %+v\n", req)
	var (
		response = &models.UserGetListResponse{}
		offset   = c.Cfg.DefaultOffset
		limit    = c.Cfg.DefaultLimit
	)

	if req.Offset > 0 {
		offset = req.Offset
	}

	if req.Limit > 0 {
		limit = req.Limit
	}

	response.Count = len(c.Users)
	if len(c.Users) < limit+offset {
		response.Users = c.Users[offset:]
	} else {
		response.Users = c.Users[offset : offset+limit]
	}

	return response
}

func (c *Controller) GetList(search string) *models.UserGetListResponse {
	log.Printf("UserGetSearch for: %+v\n", search)
	var (
		response = &models.UserGetListResponse{}
		// searchName = &models.UserGetListResponse{}
	)

	// response.Count = len(c.Users)
	// if len(c.Users) < limit+offset {
	// 	response.Users = c.Users[offset:]
	// 	for _, user := range response.Users {
	// 		if user.Name == search {
	// 			searchName.Users = append(searchName.Users, user)
	// 		}
	// 	}
	// } else {
	// 	response.Users = c.Users[offset : offset+limit]
	// 	for _, user := range response.Users {
	// 		if user.Name == search {
	// 			searchName.Users = append(searchName.Users, user)
	// 		}
	// 	}
	// }
	for _, user := range c.Users {
		if strings.Contains(user.Name, search) {
			response.Users = append(response.Users, user)
			response.Count++
		}
	}

	return response

}
