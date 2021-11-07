package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/ryokubozono/go-docker/pkg/e"
	"github.com/ryokubozono/go-docker/pkg/util"
)

type tokenRequest struct {
	Token string
}

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var request tokenRequest

		if err := c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}

		code = e.SUCCESS
		token := request.Token
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
