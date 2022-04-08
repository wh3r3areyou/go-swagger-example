package parameters

import (
	"github.com/example/example/pkg/models"
)

type AddPetParams struct {

	/* Body.

	   Pet object that needs to be added to the store
	*/

	Body models.Pet
}
