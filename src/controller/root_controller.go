package controller

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/ryokubozono/go-docker/entity"
    "github.com/ryokubozono/go-docker/pkg/util"
    "github.com/ryokubozono/go-docker/service"
)

type RootController struct{}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Success 200
// @Router / [get]
func (pc RootController) RootGet(c *gin.Context) {
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
func (pc RootController) DbPing(c *gin.Context) {
    log.Print("call DbPing")
    var s service.RootService
    user, _ := s.FirstSampleTable(c)

    c.JSON(200, gin.H{
        "hello": user.Name,
    })

}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Param username body entity.CreateUserRequest true "CreateUserRequest"
// @Success 200
// @Router /create_user [post]
func (pc RootController) CreateUser(c *gin.Context) {

    var request entity.CreateUserRequest

    if err := c.BindJSON(&request); err != nil {
        c.String(http.StatusBadRequest, "Bad request")
        return
    } else {
        var s service.RootService
        err := s.CreateUser(request.Username, request.Password)
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
        c.Status(http.StatusOK)
        return
    }
}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Param username body entity.LoginRequest true "LoginRequest"
// @Success 200
// @Router /login [post]
func (pc RootController) Login(c *gin.Context) {

    var request entity.LoginRequest

    if err := c.BindJSON(&request); err != nil {
        c.String(http.StatusBadRequest, "Bad request")
        return
    } else {
        var s service.RootService
        if err := s.Login(request.Username, request.Password); err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        } else {
            token, err := util.GenerateToken(request.Username, request.Password)
            if err != nil {
                c.String(http.StatusBadRequest, "Bad request")
                return
            }
            c.JSON(http.StatusOK, gin.H{
                "token": token,
            })
            return
        }
    }
}

// @Tags Root
// @Accept json
// @Produce json
// @accept application/x-json-stream
// @Success 200
// @Router /logout [get]
func (pc RootController) Logout(c *gin.Context) {

}
