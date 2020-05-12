package controllers

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/andrushin-anton/golang-microservices/mvc/services"
	"github.com/andrushin-anton/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request", 
		}
		// return bad request 
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// handle the error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return 
	}

	// return User to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}