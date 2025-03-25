package apicontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (controller *Controller) ok(c *gin.Context, data any) {

	c.JSON(http.StatusOK, data)
}

func (controller *Controller) badRequest(c *gin.Context, msg string) {

	c.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})
}

func (controller *Controller) error(c *gin.Context, err error) {

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
