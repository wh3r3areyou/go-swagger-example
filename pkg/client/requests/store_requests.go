package requests

type StoreRequests struct {
}

type (
	Storeer interface {
	}
)

func InitStoreRequests() Storeer {
	return &StoreRequests{}
}
