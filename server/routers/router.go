package routers

import (
	"fmt"
	"time"

	routers "github.com/flutter_go/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	tags := r.Group("/tag")
	{
		tags.GET("/:key", routers.ListTags)
	}

	users := r.Group("/user")
	{
		users.POST("/", routers.CreateUser)
		users.POST("/upload", routers.UploadUsers)
		users.POST("/login", routers.LoginUser)
		users.GET("/logout", routers.LogoutUser)
		// users.GET("/", routers.ListUsers) //This is not allowed here. User list can be via the categories API
		users.GET("/:username", routers.GetUserByName)
		users.PUT("/:username", routers.UpdateUser)
		users.DELETE("/:username", routers.DeleteUser)
	}

	apis := r.Group("/apis")
	{
		apis.POST("/", routers.AddAPI)
		apis.GET("/", routers.ListAPIs)
		apis.GET("/:apiName", routers.ListChildrenAPIs)
		apis.GET("/:apiName/:itemId/:apiN", routers.ListChildItems)
	}

	items := r.Group("/api/:apiName")
	{
		items.POST("/", routers.AddItem)
		items.POST("/uploadItems", routers.UploadItems)
		items.POST("/:itemId/uploadFile", routers.UploadFile)
		items.GET("/findByStatus", routers.FindItemsByStatus)
		items.GET("/findByTags", routers.FindItemsByTags)
		items.GET("/:itemId", routers.GetItemById)
		items.PUT("/:itemId", routers.UpdateItem)
		items.DELETE("/:itemId", routers.DeleteItem)
	}
	return r
}
