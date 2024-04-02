package handler

import (
	"net/http"

	"github.com/Daniel-Fonseca-da-Silva/Goppotunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description list a jobs openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
