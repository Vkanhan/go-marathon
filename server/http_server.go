package server

import (
	"database/sql"
	"log"

	"github.com/Vkanhan/go-marathon/controllers"
	"github.com/Vkanhan/go-marathon/repositories"
	"github.com/Vkanhan/go-marathon/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultsController
}

// InitHttpServer initializes and configures the HTTP server.
func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	runnersRepository := repositories.NewRunnersRepository(dbHandler)
	resultRepository := repositories.NewResultsRepository(dbHandler)

	runnersService := services.NewRunnersService(runnersRepository, resultRepository)
	resultsService := services.NewResultsService(resultRepository, runnersRepository)

	runnersController := controllers.NewRunnerController(runnersService)
	resultsController := controllers.NewResultsController(resultsService)

	router := gin.Default()

	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)

	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result/:id", resultsController.DeleteResult)

	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
	}
}

func (hs *HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("error starting http server: %v", err)
	}
}
