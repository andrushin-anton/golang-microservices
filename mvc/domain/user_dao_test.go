package domain

import (
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when a user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestGetNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "we were not expecting an error")
	assert.NotNil(t, user, "we were expecting an actual user")
	assert.EqualValues(t, user.Id, 123)
	assert.EqualValues(t, user.FirstName, "Anton")
	assert.EqualValues(t, user.LastName, "Andrushin")
	assert.EqualValues(t, user.Email, "andrushin.anton@gmail.com")
}