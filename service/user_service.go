package service

import (
	"Easy_IM/models"

	"github.com/gin-gonic/gin"
)

/* shift alt a */

// ctrl /

// GetUserList
// @Tags 用户列表
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
