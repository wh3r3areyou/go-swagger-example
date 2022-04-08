package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/example/example/internal/controllers"
	"github.com/example/example/internal/handler"
	"github.com/example/example/internal/repositories"
	"github.com/example/example/internal/requests"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run() {

	db, err := initDB()

	if err != nil {
		logrus.Fatal("Error with create DB")
	}

	repos := repositories.InitRepositories(db)
	requests := requests.InitRequests()
	controllers := controllers.InitControllers(repos, requests)

	server := handler.InitHandler(controllers)
	server.RunServer()
}

func initDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"postgresql", "5432", "postgres", "postgres", os.Getenv("POSTGRES_PASS")))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}
