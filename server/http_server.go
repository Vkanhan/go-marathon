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

// HttpServer represents the structure for the HTTP server
type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultsController
	usersController   *controllers.UsersController
}

// InitHttpServer initializes and configures the HTTP server from lower layer to upward.
func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	// Initialize repositories
	runnersRepository := repositories.NewRunnersRepository(dbHandler)
	resultRepository := repositories.NewResultsRepository(dbHandler)
	usersRepository := repositories.NewUsersRepository(dbHandler)

	// Initialize services
	runnersService := services.NewRunnersService(runnersRepository, resultRepository)
	resultsService := services.NewResultsService(resultRepository, runnersRepository)
	usersService := services.NewUsersService(usersRepository)

	// Initialize controllers
	runnersController := controllers.NewRunnersController(runnersService, usersService)
	resultsController := controllers.NewResultsController(resultsService, usersService)
	usersController := controllers.NewUsersController(usersService)

	router := gin.Default()

	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)

	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result/:id", resultsController.DeleteResult)

	router.POST("/login", usersController.Login)
	router.POST("logout", usersController.Logout)

	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
		usersController:   usersController,
	}
}

// Start runs the HTTP server using the configured address
func (hs *HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("error starting http server: %v", err)
	}
}
