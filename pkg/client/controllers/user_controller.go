package controllers

import (
	"github.com/example/example/internal/response"
	"github.com/example/example/pkg/client/repositories"
	"github.com/example/example/pkg/client/requests"
	"github.com/example/example/pkg/models/parameters"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	repo     repositories.Userer
	requests requests.Userer
}

type (
	Userer interface {
		CreateUser

		CreateUsersWithArrayInput

		CreateUsersWithListInput

		DeleteUser

		GetUserByName

		LoginUser

		LogoutUser

		UpdateUser
	}

	CreateUser interface {
		CreateUser(c *gin.Context)
	}

	CreateUsersWithArrayInput interface {
		CreateUsersWithArrayInput(c *gin.Context)
	}

	CreateUsersWithListInput interface {
		CreateUsersWithListInput(c *gin.Context)
	}

	DeleteUser interface {
		DeleteUser(c *gin.Context)
	}

	GetUserByName interface {
		GetUserByName(c *gin.Context)
	}

	LoginUser interface {
		LoginUser(c *gin.Context)
	}

	LogoutUser interface {
		LogoutUser(c *gin.Context)
	}

	UpdateUser interface {
		UpdateUser(c *gin.Context)
	}
)

func InitUserController(repo repositories.Userer, requests requests.Userer) Userer {
	return &UserController{
		repo:     repo,
		requests: requests,
	}
}

// CreateUser CreateUser godoc
// @Summary Create user
// @ID CreateUser
// @Schemes
// @Description This can only be done by the logged in user.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user [POST]
func (s *UserController) CreateUser(c *gin.Context) {

	var params parameters.CreateUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// CreateUsersWithArrayInput CreateUsersWithArrayInput godoc
// @Summary Creates list of users with given input array
// @ID CreateUsersWithArrayInput
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/createWithArray [POST]
func (s *UserController) CreateUsersWithArrayInput(c *gin.Context) {

	var params parameters.CreateUsersWithArrayInputParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// CreateUsersWithListInput CreateUsersWithListInput godoc
// @Summary Creates list of users with given input array
// @ID CreateUsersWithListInput
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/createWithList [POST]
func (s *UserController) CreateUsersWithListInput(c *gin.Context) {

	var params parameters.CreateUsersWithListInputParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// DeleteUser DeleteUser godoc
// @Summary Delete user
// @ID DeleteUser
// @Schemes
// @Description This can only be done by the logged in user.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/{username} [DELETE]
func (s *UserController) DeleteUser(c *gin.Context) {

	var params parameters.DeleteUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// GetUserByName GetUserByName godoc
// @Summary Get user by user name
// @ID GetUserByName
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/{username} [GET]
func (s *UserController) GetUserByName(c *gin.Context) {

	username := c.Param("username")
	logrus.Println(username)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// LoginUser LoginUser godoc
// @Summary Logs user into the system
// @ID LoginUser
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/login [GET]
func (s *UserController) LoginUser(c *gin.Context) {

	password := c.Param("password")
	logrus.Println(password)

	username := c.Param("username")
	logrus.Println(username)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// LogoutUser LogoutUser godoc
// @Summary Logs out current logged in user session
// @ID LogoutUser
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/logout [GET]
func (s *UserController) LogoutUser(c *gin.Context) {

	response.NewSuccessResponse(c, 200, struct{}{})

}

// UpdateUser UpdateUser godoc
// @Summary Updated user
// @ID UpdateUser
// @Schemes
// @Description This can only be done by the logged in user.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /user/{username} [PUT]
func (s *UserController) UpdateUser(c *gin.Context) {

	var params parameters.UpdateUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}
