package routers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	postgrefuncs "github.com/flutter_go/app/database/postgres"
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
	id := ctx.Param("username")
	post := postgrefuncs.FetchUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"username": post.UserName,
	})
}

func CreateUser(ctx *gin.Context) {
	//TODO: add validations
	u := ctx.Query("username")
	f := ctx.Query("firstname")
	l := ctx.Query("lastname")
	e := ctx.Query("email")
	p := ctx.Query("password") //assume we are receiving hashed pwd only from client. TODO add more security
	ph := ctx.Query("phone")

	id := postgrefuncs.AddUser(u, f, l, e, p, ph)

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
