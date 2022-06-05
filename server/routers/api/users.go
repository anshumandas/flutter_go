package routers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	postgrefuncs "github.com/flutter_go/database/postgres/funcs"
	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {

	posts := postgrefuncs.FetchUsers()

	ctx.JSON(http.StatusOK, posts)
}

func LogoutUser(ctx *gin.Context) {

	posts := postgrefuncs.FetchUsers()

	ctx.JSON(http.StatusOK, posts)
}

func GetUserByName(ctx *gin.Context) {
	post := postgrefuncs.FetchLastUser()

	ctx.JSON(http.StatusOK, gin.H{
		"name": post.Name,
	})
}

func CreateUser(ctx *gin.Context) {

	t := ctx.Query("title")
	d := ctx.Query("description")
	a := ctx.Query("tags")

	id := postgrefuncs.AddUser(t, d, strings.Split(a, ","))

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func UploadUsers(ctx *gin.Context) {
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	cnvid, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Failed to delete a post: %v\n", err)
	}
	postgrefuncs.DeleteUser(cnvid)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	t := ctx.Query("name")
	d := ctx.Query("description")
	a := ctx.Query("Users")

	i := strings.ReplaceAll(id, " ", "")
	cnvid, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("Failed to update a post: %v\n", err)
	}

	postgrefuncs.UpdateUser(cnvid, t, d, a)

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "updated",
		"id":          cnvid,
		"name":        t,
		"description": d,
		"tags":        a,
	})
}
