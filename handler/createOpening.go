package handler

import (
	"net/http"

	"github.com/Daniel-Fonseca-da-Silva/Goppotunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary create opening
// @Description Create a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opining: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro creating opening on database")
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
