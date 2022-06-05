package routers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	redisfuncs "github.com/flutter_go/database/redis/funcs"
	"github.com/flutter_go/models"
	"github.com/gin-gonic/gin"
)

func AddItem(ctx *gin.Context) {

	var (
		api       = ctx.Param("apiName")
		name      = ctx.Query("name")
		desc      = ctx.Query("description")
		tags      = ctx.Query("tags")
		createdBy = ctx.Query("createdBy")         //TODO this is from the user login context
		key       = fmt.Sprintf("%s%s", api, name) //This can be UUID if name is not the primary key
	)

	err := redisfuncs.CreateItem(key, name, desc, strings.Split(tags, ","), createdBy)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Item with this title already exists!",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":   name,
			"author": createdBy,
			"key":    key,
		})
	}
}

func UpdateItem(ctx *gin.Context) {
}

func UploadItems(ctx *gin.Context) {
}

func UploadFile(ctx *gin.Context) {
}

func FindItemsByStatus(ctx *gin.Context) {
}

func FindItemsByTags(ctx *gin.Context) {
}

func GetItemById(ctx *gin.Context) {
	var (
		err  error
		a    models.Item
		keys []string
	)

	key := ctx.Param("key")
	all := ctx.Query("all")

	if all != "true" {
		a, _, err = redisfuncs.GetItem(key, false)
	} else {
		a, keys, err = redisfuncs.GetItem(key, true)
	}

	if err != nil {
		noMatch(ctx)
	} else if all != "true" {
		ctx.JSON(http.StatusOK, gin.H{
			"key":       key,
			"title":     a.Name,
			"createdBy": a.CreatedBy,
		})
	} else {
		var anns = []models.Item{}

		for _, k := range keys {
			a, _, err = redisfuncs.GetItem(k, false)
			anns = append(anns, a)
		}

		ctx.JSON(http.StatusOK, anns)

	}
}

func DeleteItem(ctx *gin.Context) {
	key := ctx.Param("key")

	err := redisfuncs.Delete(key)
	if err != nil {
		log.Printf("Failed to delete item: %v", err)
	}
}

func noItemMatch(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "There is no item with this title",
	})
}
