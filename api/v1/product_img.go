package v1

import (
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func ListProductsImg(c *gin.Context) {
	listProductImg := services.ListProductImg{}
	if err := c.ShouldBind(&listProductImg); err == nil {
		res := listProductImg.ListImg(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
