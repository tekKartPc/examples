package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	iris.Use(customLogger)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	iris.Get("/1", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	iris.Get("/2", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	// log http errors
	errorLogger := logger.New()

	// yes we have options look at the logger.Options inside kataras/iris/middleware/logger.go
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("My Custom 404 error page ")
	})
	//

	iris.Listen(":8080")

}

/* Book section: 'Logger' */
