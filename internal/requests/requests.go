package requests

import (
	"github.com/example/example/pkg/client/requests"
)

type Requests struct {
	PetRequests requests.Peter

	StoreRequests requests.Storeer

	UserRequests requests.Userer
}

func InitRequests() *Requests {
	return &Requests{

		PetRequests: requests.InitPetRequests(),

		StoreRequests: requests.InitStoreRequests(),

		UserRequests: requests.InitUserRequests(),
	}
}
