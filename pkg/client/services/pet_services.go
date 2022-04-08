package services

type PetService struct {
}

type (
	Peter interface {
	}
)

func InitPetService() Peter {
	return &PetService{}
}
