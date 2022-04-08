package parameters

type UpdatePetWithFormParams struct {

	/* Name.

	   Updated name of the pet
	*/

	Name string

	/* PetID.

	   ID of pet that needs to be updated

	   Format: int64
	*/

	PetID int64

	/* Status.

	   Updated status of the pet
	*/

	Status string
}
