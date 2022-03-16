package routers

import (
	_ "blog-service/docs"
	"github.com/gin-gonic/gin"
	"blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 设计完成后，进行基础编码，确定好方法原型
func NewRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//url := ginSwagger.URL("http://127.0.0.1:8010/swagger/doc.json")		//手动指定swagger.json的路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))		// 初始化 docs 包和注册一个针对 swagger 的路由
	
	article := v1.NewArticle()
	tag := v1.NewTag()
	// 有处理Handler方法后，注册到对应的路由规则上
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}