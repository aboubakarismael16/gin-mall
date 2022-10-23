package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	createCartService := services.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowCarts 购物车详细信息
func ShowCarts(c *gin.Context) {
	showCartsService := services.CartService{}
	res := showCartsService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// UpdateCart 修改购物车信息
func UpdateCart(c *gin.Context) {
	updateCartService := services.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	deleteCartService := services.CartService{}
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}