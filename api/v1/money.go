package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func ShowMoney(c *gin.Context) {
	showMoneyService := services.ShowMoneyService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err == nil {
		res := showMoneyService.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
