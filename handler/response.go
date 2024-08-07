package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string){
	ctx.JSON(code, gin.H{
		"message": msg,
		"status": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data": data,
	})
}

type ErrorResponse struct{
	Message string `json:"message"`
	Status string `json:"status"`
}

type OpeningResponse struct{
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct{
	Message string `json:"message"`
	Data []schemas.OpeningResponse `json:"data"`
}