package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/service"
	"github.com/ryokubozono/go-docker/entity"
    "log"
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
    log.Print("call DbPing")
    var s root.Service
    user, _ := s.FirstSampleTable(c)

    c.JSON(200, gin.H{
        "hello": user.Name,
    })

}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Param username body entity.UsernamePasswordSignUpRequest true "UsernamePasswordSignUpRequest"
// @Success 200
// @Router /signup [post]
func (pc Controller) SignUp(c *gin.Context) {

    var request entity.UsernamePasswordSignUpRequest

    err := c.BindJSON(&request)

    if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
        return    
    } else {
        var s root.Service
        err := s.UsernamePasswordSignUp(request.Username, request.Password)
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
    }

    c.Status(http.StatusOK)
    return
}