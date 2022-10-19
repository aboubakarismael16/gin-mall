package v1

import (
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func ListCarousels(c *gin.Context) {
	listCarouselsService := services.ListCarouselsService{}
	if err := c.ShouldBind(&listCarouselsService); err == nil {
		res := listCarouselsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
