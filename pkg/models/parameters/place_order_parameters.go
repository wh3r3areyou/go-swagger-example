package parameters

import (
	"github.com/example/example/pkg/models"
)

type PlaceOrderParams struct {

	/* Body.

	   order placed for purchasing the pet
	*/

	Body models.Order
}
