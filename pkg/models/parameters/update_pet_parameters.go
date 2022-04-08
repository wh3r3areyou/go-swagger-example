package parameters

import (
	"github.com/example/example/pkg/models"
)

type UpdatePetParams struct {

	/* Body.

	   Pet object that needs to be added to the store
	*/

	Body models.Pet
}
