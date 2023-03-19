package routes

import (
	"ginblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IninRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //初始化路由

	router := r.Group("api/v1") //设置路由组
	{
		router.GET("hellow", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
