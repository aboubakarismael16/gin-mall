package v1

import (
	util "gin-mall/pkg/utils"
	"gin-mall/services"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegisterService services.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLoginService services.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdateService services.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.Update(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadAvatarService := services.UserService{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err == nil {
		res := uploadAvatarService.Post(c.Request.Context(), chaim.ID, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmailService services.SendEmailService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmailService); err == nil {
		res := sendEmailService.Send(c.Request.Context(), chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func ValidEmail(c *gin.Context) {
	var vaildEmailService services.ValidEmailService
	if err := c.ShouldBind(vaildEmailService); err == nil {
		res := vaildEmailService.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
