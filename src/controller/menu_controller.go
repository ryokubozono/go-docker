package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

type MenuController struct{}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Success 200
// @Security BasicAuth
// @Router /menu/top [get]    
func (pc MenuController) Top(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "hello world",
    })
}