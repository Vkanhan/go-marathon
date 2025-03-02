package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Vkanhan/go-marathon/models"
	"github.com/Vkanhan/go-marathon/services"
	"github.com/gin-gonic/gin"
)

type RunnersController struct {
	runnersService *services.RunnersService
	usersService   *services.UsersService
}

func NewRunnersController(runnersService *services.RunnersService, usersService *services.UsersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
		usersService:   usersService,
	}
}

// CreateRunner handles the creation of a new runner.
func (rh *RunnersController) CreateRunner(ctx *gin.Context) {
	// Authentication and authorization check
	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ADMIN_ROLE})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := rh.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// UpdateRunner handles the update of an existing runner.
func (rh *RunnersController) UpdateRunner(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ADMIN_ROLE})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var runner models.Runner
	if err := json.Unmarshal(body, &runner); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr = rh.runnersService.UpdateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)

}

// DeleteRunner handles the deletion of a runner by ID.2
func (rh *RunnersController) DeleteRunner(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ADMIN_ROLE})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	runnerID := ctx.Param("id")
	responseErr = rh.runnersService.DeleteRunner(runnerID)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetRunner handles retrieving a runner by ID.
func (rh *RunnersController) GetRunner(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ADMIN_ROLE, RUNNER_ROLE})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	runnerID := ctx.Param("id")
	response, responseErr := rh.runnersService.GetRunner(runnerID)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetRunnersBatch handles retrieving a batch of runners.
func (rh *RunnersController) GetRunnersBatch(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ADMIN_ROLE, RUNNER_ROLE})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")

	response, responseErr := rh.runnersService.GetRunnersBatch(country, year)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
