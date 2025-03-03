package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Vkanhan/go-marathon/models"
	"github.com/Vkanhan/go-marathon/services"
	"github.com/gin-gonic/gin"
)

const ADMIN_ROLE = "admin"
const RUNNER_ROLE = "runner"

// ResultsController manages result-related requests.
type ResultsController struct {
	resultsService *services.ResultsService
	usersService   *services.UsersService
}

func NewResultsController(resultsService *services.ResultsService, usersService *services.UsersService) *ResultsController {
	return &ResultsController{
		resultsService: resultsService,
		usersService:   usersService,
	}
}

func (rh *ResultsController) CreateResult(ctx *gin.Context) {
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
	var results models.Result
	if err := json.Unmarshal(body, &results); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := rh.resultsService.CreateResult(&results)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (rh *ResultsController) DeleteResult(ctx *gin.Context) {
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

	resultID := ctx.Param("id")
	responseErr = rh.resultsService.DeleteResult(resultID)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
