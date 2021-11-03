package root

import (
	"github.com/gin-gonic/gin"

	"github.com/ryokubozono/go-docker/entity"
    "github.com/ryokubozono/go-docker/db"

)

type TestTable entity.TestTable
type Service struct{}

func (s Service) FirstTestTable(c *gin.Context) (TestTable, error){

	db := db.GetDB()
	var user TestTable

	db.First(&user, 1)

	return user, nil
}
