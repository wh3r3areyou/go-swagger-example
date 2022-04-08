package controllers

import (
	"github.com/example/example/internal/response"
	"github.com/example/example/pkg/client/repositories"
	"github.com/example/example/pkg/client/requests"
	"github.com/example/example/pkg/models/parameters"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PetController struct {
	repo     repositories.Peter
	requests requests.Peter
}

type (
	Peter interface {
		AddPet

		DeletePet

		FindPetsByStatus

		FindPetsByTags

		GetPetByID

		UpdatePet

		UpdatePetWithForm

		UploadFile
	}

	AddPet interface {
		AddPet(c *gin.Context)
	}

	DeletePet interface {
		DeletePet(c *gin.Context)
	}

	FindPetsByStatus interface {
		FindPetsByStatus(c *gin.Context)
	}

	FindPetsByTags interface {
		FindPetsByTags(c *gin.Context)
	}

	GetPetByID interface {
		GetPetByID(c *gin.Context)
	}

	UpdatePet interface {
		UpdatePet(c *gin.Context)
	}

	UpdatePetWithForm interface {
		UpdatePetWithForm(c *gin.Context)
	}

	UploadFile interface {
		UploadFile(c *gin.Context)
	}
)

func InitPetController(repo repositories.Peter, requests requests.Peter) Peter {
	return &PetController{
		repo:     repo,
		requests: requests,
	}
}

// AddPet AddPet godoc
// @Summary Add a new pet to the store
// @ID AddPet
// @Schemes
// @Description
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet [POST]
func (s *PetController) AddPet(c *gin.Context) {

	var params parameters.AddPetParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// DeletePet DeletePet godoc
// @Summary Deletes a pet
// @ID DeletePet
// @Schemes
// @Description
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/{petId} [DELETE]
func (s *PetController) DeletePet(c *gin.Context) {

	var params parameters.DeletePetParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// FindPetsByStatus FindPetsByStatus godoc
// @Summary Finds Pets by status
// @ID FindPetsByStatus
// @Schemes
// @Description Multiple status values can be provided with comma separated strings
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/findByStatus [GET]
func (s *PetController) FindPetsByStatus(c *gin.Context) {

	status := c.Param("status")
	logrus.Println(status)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// FindPetsByTags FindPetsByTags godoc
// @Summary Finds Pets by tags
// @ID FindPetsByTags
// @Schemes
// @Description Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/findByTags [GET]
func (s *PetController) FindPetsByTags(c *gin.Context) {

	tags := c.Param("tags")
	logrus.Println(tags)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// GetPetByID GetPetByID godoc
// @Summary Find pet by ID
// @ID GetPetByID
// @Schemes
// @Description Returns a single pet
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/{petId} [GET]
func (s *PetController) GetPetByID(c *gin.Context) {

	petId := c.Param("petId")
	logrus.Println(petId)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// UpdatePet UpdatePet godoc
// @Summary Update an existing pet
// @ID UpdatePet
// @Schemes
// @Description
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet [PUT]
func (s *PetController) UpdatePet(c *gin.Context) {

	var params parameters.UpdatePetParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// UpdatePetWithForm UpdatePetWithForm godoc
// @Summary Updates a pet in the store with form data
// @ID UpdatePetWithForm
// @Schemes
// @Description
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/{petId} [POST]
func (s *PetController) UpdatePetWithForm(c *gin.Context) {

	var params parameters.UpdatePetWithFormParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// UploadFile UploadFile godoc
// @Summary uploads an image
// @ID UploadFile
// @Schemes
// @Description
// @Tags pet
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /pet/{petId}/uploadImage [POST]
func (s *PetController) UploadFile(c *gin.Context) {

	var params parameters.UploadFileParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}
