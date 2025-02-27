package handler

import (
	"net/http"

	"github.com/NatalNW7/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

//  @BasePath /api/v1

// @Sumary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context){
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "listing-openings", openings)
}