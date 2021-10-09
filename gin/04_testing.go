package main

import (
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/restuwahyu13/go-testing-fundamental/gin/utils"
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
	body, _ := io.ReadAll(ctx.Request.Body)

	if reflect.DeepEqual(len(string(body)), int(2)) {
		message := Response{
			Code:    http.StatusOK,
			Message: "Request From Post Method",
		}
		ctx.JSON(http.StatusOK, message)
	} else {
		parse := utils.Parse(body)
		message := Response{
			Code:    http.StatusOK,
			Message: parse.Message,
		}
		ctx.JSON(http.StatusOK, message)
	}

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
