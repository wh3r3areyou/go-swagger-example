package models

import "github.com/go-openapi/strfmt"

// Order order
//
// swagger:model Order
type Order struct {

	// complete
	Complete bool `json:"complete,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// pet Id
	PetID int64 `json:"petId,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// ship date
	// Format: date-time
	ShipDate strfmt.DateTime `json:"shipDate,omitempty"`

	// Order Status
	// Enum: [placed approved delivered]
	Status string `json:"status,omitempty"`
}
