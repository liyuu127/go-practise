package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liyuu127/go-practise/swagger/api"
	_ "github.com/liyuu127/go-practise/swagger/docs"
	"net/http"
)

var users []api.User

func main() {

	r := gin.Default()
	r.POST("/users", Create)
	r.GET("/users/:username", Get)

}

func Create(ctx *gin.Context) {

	var user api.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}

	users = append(users, user)
	ctx.JSON(http.StatusOK, users)

}

func Get(ctx *gin.Context) {
	username := ctx.Param("username")

	for _, u := range users {
		if u.Name == username {
			ctx.JSON(http.StatusOK, u)
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})

}
