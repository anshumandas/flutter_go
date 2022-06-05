package routers

import (
	"net/http"

	postgrefuncs "github.com/flutter_go/database/postgres/funcs"
	"github.com/gin-gonic/gin"
)

func ListTags(ctx *gin.Context) {

	posts := postgrefuncs.FetchTags()

	ctx.JSON(http.StatusOK, posts)
}
