package main

import (
	"github.com/kataras/iris/v12"
)

func index(c iris.Context) {
	c.JSON(iris.Map{
		"code": 200,
		"msg":  "success",
	})
}

func main() {
	app := iris.New()
	app.Get("/", index)
	app.Listen(":8080")
}
