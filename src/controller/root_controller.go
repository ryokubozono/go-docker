package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/service"

)

type Controller struct{}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Success 200
// @Router / [get]    
func (pc Controller) RootGet(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "hello world",
    })
}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Success 200
// @Router /ping [get]
func (pc Controller) DbPing(c *gin.Context) {

    var s root.Service
    user, _ := s.FirstTestTable(c)

    c.JSON(200, gin.H{
        "hello": user.Name,
    })

}