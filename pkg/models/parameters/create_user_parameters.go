package parameters

import (
	"github.com/example/example/pkg/models"
)

type CreateUserParams struct {

	/* Body.

	   Created user object
	*/

	Body models.User
}
