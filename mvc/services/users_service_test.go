package services

import (
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/andrushin-anton/golang-microservices/mvc/utils"
	"github.com/andrushin-anton/golang-microservices/mvc/domain"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError) 
)

func inti() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDataBase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 does not exists",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Message, "user 0 not found")
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
			FirstName: "Anton",
			LastName: "Andrushin",
			Email: "andrushin.anton@gmail.com",
		}, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 123)
	assert.EqualValues(t, user.FirstName, "Anton")
}