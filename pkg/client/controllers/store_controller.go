package controllers

import (
	"github.com/example/example/internal/response"
	"github.com/example/example/pkg/client/repositories"
	"github.com/example/example/pkg/client/requests"
	"github.com/example/example/pkg/models/parameters"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StoreController struct {
	repo     repositories.Storeer
	requests requests.Storeer
}

type (
	Storeer interface {
		DeleteOrder

		GetInventory

		GetOrderByID

		PlaceOrder
	}

	DeleteOrder interface {
		DeleteOrder(c *gin.Context)
	}

	GetInventory interface {
		GetInventory(c *gin.Context)
	}

	GetOrderByID interface {
		GetOrderByID(c *gin.Context)
	}

	PlaceOrder interface {
		PlaceOrder(c *gin.Context)
	}
)

func InitStoreController(repo repositories.Storeer, requests requests.Storeer) Storeer {
	return &StoreController{
		repo:     repo,
		requests: requests,
	}
}

// DeleteOrder DeleteOrder godoc
// @Summary Delete purchase order by ID
// @ID DeleteOrder
// @Schemes
// @Description For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /store/order/{orderId} [DELETE]
func (s *StoreController) DeleteOrder(c *gin.Context) {

	var params parameters.DeleteOrderParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}

// GetInventory GetInventory godoc
// @Summary Returns pet inventories by status
// @ID GetInventory
// @Schemes
// @Description Returns a map of status codes to quantities
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /store/inventory [GET]
func (s *StoreController) GetInventory(c *gin.Context) {

	response.NewSuccessResponse(c, 200, struct{}{})

}

// GetOrderByID GetOrderByID godoc
// @Summary Find purchase order by ID
// @ID GetOrderByID
// @Schemes
// @Description For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /store/order/{orderId} [GET]
func (s *StoreController) GetOrderByID(c *gin.Context) {

	orderId := c.Param("orderId")
	logrus.Println(orderId)

	response.NewSuccessResponse(c, 200, struct{}{})

}

// PlaceOrder PlaceOrder godoc
// @Summary Place an order for a pet
// @ID PlaceOrder
// @Schemes
// @Description
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {object} response.successResponse
// @Router /store/order [POST]
func (s *StoreController) PlaceOrder(c *gin.Context) {

	var params parameters.PlaceOrderParams

	if err := c.ShouldBindJSON(&params); err != nil {
		logrus.Println(err.Error())
		response.NewErrorResponse(c, 400, err.Error())
		return
	}

	response.NewSuccessResponse(c, 200, struct{}{})

}
