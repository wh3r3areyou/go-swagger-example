package handler

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	docs "github.com/example/example/docs"
	"github.com/example/example/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Controllers *controllers.Controllers
}

func InitHandler(controllers *controllers.Controllers) *Handler {
	return &Handler{Controllers: controllers}
}

func (h *Handler) initRoutes(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/v2")

	v1.POST(ginizePath("/pet"), h.Controllers.PetController.AddPet)

	v1.POST(ginizePath("/user"), h.Controllers.UserController.CreateUser)

	v1.POST(ginizePath("/user/createWithArray"), h.Controllers.UserController.CreateUsersWithArrayInput)

	v1.POST(ginizePath("/user/createWithList"), h.Controllers.UserController.CreateUsersWithListInput)

	v1.DELETE(ginizePath("/store/order/{orderId}"), h.Controllers.StoreController.DeleteOrder)

	v1.DELETE(ginizePath("/pet/{petId}"), h.Controllers.PetController.DeletePet)

	v1.DELETE(ginizePath("/user/{username}"), h.Controllers.UserController.DeleteUser)

	v1.GET(ginizePath("/pet/findByStatus"), h.Controllers.PetController.FindPetsByStatus)

	v1.GET(ginizePath("/pet/findByTags"), h.Controllers.PetController.FindPetsByTags)

	v1.GET(ginizePath("/store/inventory"), h.Controllers.StoreController.GetInventory)

	v1.GET(ginizePath("/store/order/{orderId}"), h.Controllers.StoreController.GetOrderByID)

	v1.GET(ginizePath("/pet/{petId}"), h.Controllers.PetController.GetPetByID)

	v1.GET(ginizePath("/user/{username}"), h.Controllers.UserController.GetUserByName)

	v1.GET(ginizePath("/user/login"), h.Controllers.UserController.LoginUser)

	v1.GET(ginizePath("/user/logout"), h.Controllers.UserController.LogoutUser)

	v1.POST(ginizePath("/store/order"), h.Controllers.StoreController.PlaceOrder)

	v1.PUT(ginizePath("/pet"), h.Controllers.PetController.UpdatePet)

	v1.POST(ginizePath("/pet/{petId}"), h.Controllers.PetController.UpdatePetWithForm)

	v1.PUT(ginizePath("/user/{username}"), h.Controllers.UserController.UpdateUser)

	v1.POST(ginizePath("/pet/{petId}/uploadImage"), h.Controllers.PetController.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}

func (h *Handler) RunServer() {
	router := gin.Default()
	router = h.initRoutes(router)

	var srv http.Server

	go func() {
		srv = http.Server{
			Addr:           ":" + viper.GetString("apiserver_port"),
			Handler:        router,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		}
		err := srv.ListenAndServe()

		if err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Infof("RESTApi server Example listening at %s", viper.GetString("apiserver_port"))

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("RESTApi server Example Shutting Down")
}

func ginizePath(path string) string {
	return strings.Replace(strings.Replace(path, "{", ":", -1), "}", "", -1)
}
