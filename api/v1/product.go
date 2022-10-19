package v1

import (
	"gin-mall/pkg/utils"
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := services.ProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListProducts 商品列表
func ListProducts(c *gin.Context) {
	listProductsService := services.ProductService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func SearchProduct(c *gin.Context) {
	searchProductsService := services.ProductService{}
	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.Search(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
