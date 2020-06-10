package handlers

import (
	model "library-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoHandler(c *gin.Context){
	userInfo := model.UserInfo{}
	model.DB.First(&userInfo)
	_ = c.BindJSON(&userInfo)
	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
	})
}