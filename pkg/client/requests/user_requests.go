package requests

type UserRequests struct {
}

type (
	Userer interface {
	}
)

func InitUserRequests() Userer {
	return &UserRequests{}
}
