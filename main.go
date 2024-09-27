package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
  r := gin.Default()
  r.GET("/",func(ctx *gin.Context) {
    ctx.String(200,"hello")
  })

  r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // However, this one will match /user/john/ and also /user/john/send
  // If no other routers match /user/john, it will redirect to /user/john/
  r.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  r.GET("/welcome", func(c *gin.Context) {
    c.Request.URL.Query().Get("")
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })

  r.Run()
}