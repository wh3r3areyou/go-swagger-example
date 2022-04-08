package parameters

import (
	"github.com/example/example/pkg/models"
)

type UpdateUserParams struct {

	/* Body.

	   Updated user object
	*/

	Body models.User

	/* Username.

	   name that need to be updated
	*/

	Username string
}
