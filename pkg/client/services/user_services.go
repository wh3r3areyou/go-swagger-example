package services

type UserService struct {
}

type (
	Userer interface {
	}
)

func InitUserService() Userer {
	return &UserService{}
}
