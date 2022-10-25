package routers

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})

		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProducts)
		v1.GET("imgs/:id", api.ListProductsImg)
		v1.GET("categories", api.ListCategory)

		v1.GET("carousels", api.ListCarousels) //轮播图

		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar) //上传头像
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			//add to favorite list
			authed.GET("favorites", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			//收获地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.GetAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			//购物车
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ShowCarts)  // 购物车id
			authed.PUT("carts/:id", api.UpdateCart) // 购物车id
			authed.DELETE("carts/:id", api.DeleteCart)

			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrders)
			authed.GET("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

		}
	}

	return r
}
