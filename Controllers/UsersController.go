package Controllers

import (
	"Iris-project/Models"
	"github.com/kataras/iris/v12/mvc"
	"log"
)

type UsersController struct {
}

func (this *UsersController) BeforeActivation(before mvc.BeforeActivation) {
	before.Handle("GET", "/list", "ShowAllUsers")
}

func (this *UsersController) AfterActivation(after mvc.AfterActivation) {
	log.Println("Session Done!")
}

func (this *UsersController) ShowAllUsers() map[uint64]Models.User {
	return map[uint64]Models.User{
		1: {
			ID:       1,
			UserName: "A",
			Age:      10,
		},
	}
}
