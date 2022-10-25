package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	createOrderService := services.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ListOrders(c *gin.Context) {
	listOrdersService := services.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrdersService); err == nil {
		res := listOrdersService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowOrder 订单详情
func ShowOrder(c *gin.Context) {
	showOrderService := services.OrderService{}
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteOrder(c *gin.Context) {
	deleteOrderService := services.OrderService{}
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
