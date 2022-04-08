package controllers

import (
	"github.com/example/example/internal/repositories"
	"github.com/example/example/internal/requests"
	"github.com/example/example/pkg/client/controllers"
)

type Controllers struct {
	PetController controllers.Peter

	StoreController controllers.Storeer

	UserController controllers.Userer
}

func InitControllers(repos *repositories.Repositories, requests *requests.Requests) *Controllers {
	return &Controllers{

		PetController: controllers.InitPetController(repos.PetRepository, requests.PetRequests),

		StoreController: controllers.InitStoreController(repos.StoreRepository, requests.StoreRequests),

		UserController: controllers.InitUserController(repos.UserRepository, requests.UserRequests),
	}
}
