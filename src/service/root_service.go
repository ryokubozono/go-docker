package root

import (
	"github.com/gin-gonic/gin"

	"github.com/ryokubozono/go-docker/entity"
    "github.com/ryokubozono/go-docker/db"

)

type SampleTable entity.SampleTable
type Service struct{}

func (s Service) FirstSampleTable(c *gin.Context) (SampleTable, error){

	db := db.GetDB()
	var sample SampleTable

	db.First(&sample, 1)

	return sample, nil
}
