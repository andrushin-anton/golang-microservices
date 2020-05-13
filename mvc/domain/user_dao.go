package domain

import (
	"fmt"
	"net/http"

	"github.com/andrushin-anton/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{Id:123, FirstName: "Anton", LastName: "Andrushin", Email: "andrushin.anton@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}