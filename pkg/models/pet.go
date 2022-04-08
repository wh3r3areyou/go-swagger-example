package models

// Pet pet
//
// swagger:model Pet
type Pet struct {

	// category
	Category *Category `json:"category,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: doggie
	// Required: true
	Name *string `json:"name"`

	// photo urls
	// Required: true
	PhotoUrls []string `json:"photoUrls" xml:"photoUrls"`

	// pet status in the store
	// Enum: [available pending sold]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*Tag `json:"tags" xml:"tags"`
}
