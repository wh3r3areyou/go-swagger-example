package repositories

import (
	"database/sql"

	"github.com/example/example/pkg/client/repositories"
)

type Repositories struct {
	PetRepository repositories.Peter

	StoreRepository repositories.Storeer

	UserRepository repositories.Userer
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{

		PetRepository: repositories.InitPetRepository(db),

		StoreRepository: repositories.InitStoreRepository(db),

		UserRepository: repositories.InitUserRepository(db),
	}
}
