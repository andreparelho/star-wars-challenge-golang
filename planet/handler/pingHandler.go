package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(context *gin.Context) {
	context.JSON(http.StatusOK, "Fucking pong")
}
