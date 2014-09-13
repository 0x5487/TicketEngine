package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.New()

	err := r.UseApi()
	PanicIf(err)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}
