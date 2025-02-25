package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Vkanhan/go-marathon/models"
	"github.com/Vkanhan/go-marathon/services"
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	resultsService *services.ResultsService
}

func NewResultsController(resultsService *services.ResultsService) *ResultsController {
	return &ResultsController{
		resultsService: resultsService,
	}
}

func (rh *ResultsController) CreateResult(ctx *gin.Context) {
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
	resultID := ctx.Param("id")
	responseErr := rh.resultsService.DeleteResult(resultID)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
