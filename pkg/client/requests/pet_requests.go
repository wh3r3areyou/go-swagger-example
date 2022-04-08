package requests

type PetRequests struct {
}

type (
	Peter interface {
	}
)

func InitPetRequests() Peter {
	return &PetRequests{}
}
