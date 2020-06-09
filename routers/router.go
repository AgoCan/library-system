package routers
import (
	"github.com/gin-gonic/gin"
	"library-system/middleware"
)
// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine{
	router := gin.New()
	if err := middleware.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(middleware.GinLogger(middleware.Logger),
		middleware.GinRecovery(middleware.Logger, true))

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	return router
}
