package routers

import (
	"RecordWords/service"
	"RecordWords/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找z
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", service.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", service.CreateWord)
		// 查看所有的待办事项
		v1Group.GET("/todo", service.GetWordList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", service.UpdateWord)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", service.DeleteWord)
	}
	return r
}
