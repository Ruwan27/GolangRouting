package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getData", getData)
	router.POST("/getPostData", getPostData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()

}

// hrrp://localhost:8080/getUrlData?name=Mark&age=30
func getUrlData(ctx *gin.Context) {
	name := ctx.Param(("name"))
	age := ctx.Param("age")
	ctx.JSON(200, gin.H{
		"data": "Hi I am gin framework url data",
		"name": name,
		"age":  age,
	})

}

// hrrp://localhost:8080/getQueryString?name=Mark&age=30
func getQueryString(ctx *gin.Context) {
	name := ctx.Query(("name"))
	age := ctx.Query("age")
	ctx.JSON(200, gin.H{
		"data": "Hi I am gin framework query string",
		"name": name,
		"age":  age,
	})

}
func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hi I am gin framework get method",
	})

}
func getPostData(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := ioutil.ReadAll(body)
	ctx.JSON(200, gin.H{
		"data":     "I am post method",
		"bodyData": string(value),
	})
}
