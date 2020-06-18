package handlers

import (
	"fmt"
	model "library-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoHandler(c *gin.Context) {
	var userInfos  model.UserInfo
	var blog model.Blog
	id := 2
	model.DB.Debug().Model(&blog).Where("id = ?", id).Find(&userInfos)
	fmt.Printf("%v", userInfos)
	// _ = c.BindJSON(&userInfo)
	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfos,
		"blog": blog,
	})

}
