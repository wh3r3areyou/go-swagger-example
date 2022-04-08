package services

type StoreService struct {
}

type (
	Storeer interface {
	}
)

func InitStoreService() Storeer {
	return &StoreService{}
}
