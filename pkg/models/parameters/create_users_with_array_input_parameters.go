package parameters

import (
	"github.com/example/example/pkg/models"
)

type CreateUsersWithArrayInputParams struct {

	/* Body.

	   List of user object
	*/

	Body []*models.User
}
