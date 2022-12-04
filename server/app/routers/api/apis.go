package routers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	redisfuncs "github.com/flutter_go/app/database/redis"
	"github.com/flutter_go/app/models"
	"github.com/gin-gonic/gin"
)

func AddAPI(ctx *gin.Context) {

	var (
		name     = ctx.PostForm("name")
		children = ctx.PostForm("children")
		key      = fmt.Sprintf("api%s", name)
	)

	err := redisfuncs.CreateAPI(key, name, strings.Split(children, ","))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "API with this title already exists!",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"key":  key,
		})
	}
}

func ListAPIs(ctx *gin.Context) {
	var (
		err  error
		a    models.API
		keys []string
	)

	key := ctx.Param("key")
	all := ctx.Query("all")

	if all != "true" {
		a, _, err = redisfuncs.GetAPI(key, false)
	} else {
		a, keys, err = redisfuncs.GetAPI(key, true)
	}

	if err != nil {
		noMatch(ctx)
	} else if all != "true" {
		ctx.JSON(http.StatusOK, gin.H{
			"key":  key,
			"name": a.Name,
		})
	} else {
		var anns = []models.API{}

		for _, k := range keys {
			a, _, err = redisfuncs.GetAPI(k, false)
			anns = append(anns, a)
		}

		ctx.JSON(http.StatusOK, anns)

	}
}

func ListChildrenAPIs(ctx *gin.Context) {
}

func ListChildItems(ctx *gin.Context) {
}

func DeleteAPI(ctx *gin.Context) {
	key := ctx.Param("key")

	err := redisfuncs.DeleteAPI(key)
	if err != nil {
		log.Printf("Failed to delete api: %v", err)
	}
}

func noMatch(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "There is no api with this title",
	})
}
