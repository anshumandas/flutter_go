package routers

import (
	"log"
	"net/http"
	"strings"

	redisfuncs "github.com/flutter_go/app/database/redis"
	"github.com/flutter_go/app/models"
	"github.com/gin-gonic/gin"
)

func AddItem(ctx *gin.Context) {

	var (
		api       = ctx.Param("apiName")
		name      = ctx.Query("name")
		desc      = ctx.Query("description")
		tags      = ctx.Query("tags")
		createdBy = ctx.Query("createdBy") //TODO this is from the user login context
		key       = name                   //This can be UUID if name is not the primary key
	)

	err := redisfuncs.CreateItem(api, key, name, desc, strings.Split(tags, ","), createdBy)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Item with this name already exists!",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":      name,
			"createdBy": createdBy,
			"type":      api,
			"id":        key,
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
	fetch(ctx, true, "status")
}

func FindItemsByTags(ctx *gin.Context) {
	fetch(ctx, true, "tags")
}

func GetItemById(ctx *gin.Context) {
	fetch(ctx, false)
}

func fetch(ctx *gin.Context, all bool, by ...string) {
	var (
		err  error
		a    models.Item
		keys []string
	)

	key := ctx.Param("key")
	if all {
		a, _, err = redisfuncs.GetItem(key, false)
	} else {
		a, keys, err = redisfuncs.GetItem(key, true)
		q := ctx.Query(by[0])
		if q != "" {
			log.Printf("not implemented yet")
			//cond: by=q
			//use this value to check the returned list or use RediSearch
		}
	}

	if err != nil {
		noItemMatch(ctx)
	} else if !all {
		ctx.JSON(http.StatusOK, gin.H{
			"key":       key,
			"name":      a.Name,
			"createdBy": a.CreatedBy,
		})
	} else {
		var anns = []models.Item{}

		for _, k := range keys {
			a, _, err = redisfuncs.GetItem(k, false)
			if err != nil {
				anns = append(anns, a)
			}
		}

		ctx.JSON(http.StatusOK, anns)

	}
}

func DeleteItem(ctx *gin.Context) {
	key := ctx.Param("key")

	err := redisfuncs.DeleteItem(key)
	if err != nil {
		log.Printf("Failed to delete item: %v", err)
	}
}

func noItemMatch(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "There is no item with this name",
	})
}
