package parameters

import "github.com/go-openapi/runtime"

type UploadFileParams struct {

	/* AdditionalMetadata.

	   Additional data to pass to server
	*/

	AdditionalMetadata string

	/* File.

	   file to upload
	*/

	File runtime.NamedReadCloser

	/* PetID.

	   ID of pet to update

	   Format: int64
	*/

	PetID int64
}
