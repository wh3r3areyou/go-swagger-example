package main

import (
	"github.com/example/example/internal/app"
	"github.com/example/example/internal/config"
	"github.com/joho/godotenv"
	loggrus "github.com/sirupsen/logrus"
)

// @title Swagger Petstore
// @version 1.0.6
// @license {{Apache 2.0 http://www.apache.org/licenses/LICENSE-2.0.html} {map[]}}
// @description This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
// @host petstore.swagger.io
// @BasePath /v2

func main() {
	err := config.InitConfig()

	if err != nil {
		loggrus.Fatal("Error with config = " + err.Error())
	}

	if err = godotenv.Load("/app/.env"); err != nil {
		loggrus.Fatal("Not found .env")
	}

	app.Run()
}
