package productservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service service) Fetch(c *gin.Context) {
	products, err := service.usecase.Fetch(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}
