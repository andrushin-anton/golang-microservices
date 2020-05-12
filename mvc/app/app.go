package app

import (
	"net/http"

	"github.com/andrushin-anton/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}