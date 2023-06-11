package models

type User struct {
	Name        string
	PhoneNumber string
	Birthday string
}
type UserGetListRequest struct{
	Offset int
	Limit int
}

type UserGetListResponse struct{
	Count int
	Users []*User
}