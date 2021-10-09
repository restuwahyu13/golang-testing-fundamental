package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// first function for testing

func GetRequestFirst(ctx *gin.Context) {
	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Get Method",
	}
	ctx.JSON(http.StatusOK, message)
}

func PostRequestFirst(ctx *gin.Context) {
	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Post Method",
	}
	ctx.JSON(http.StatusOK, message)
}

// second function for testing

func GetRequestSecond(ctx *gin.Context) {
	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Get Method",
	}
	ctx.JSON(http.StatusOK, message)
}

func PostRequestSecond(ctx *gin.Context) {
	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Post Method",
	}
	ctx.JSON(http.StatusOK, message)
}

// third function for testing

func GetRequestThird(ctx *gin.Context) {
	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Get Method",
	}
	ctx.JSON(http.StatusOK, message)
}

func PostRequestThird(ctx *gin.Context) {

	message := Response{
		Code:    http.StatusOK,
		Message: "Request From Post Method",
	}
	ctx.JSON(http.StatusOK, message)
}

func main() {
	app := Router()
	app.Run(":3000")
}

func Router() *gin.Engine {
	app := gin.Default()

	app.GET("/first", GetRequestFirst)
	app.POST("/first", PostRequestFirst)
	app.GET("/second", GetRequestSecond)
	app.POST("/second", PostRequestSecond)
	app.GET("/third", GetRequestThird)
	app.POST("/third", PostRequestThird)

	return app
}
