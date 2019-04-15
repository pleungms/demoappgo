package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	//http.HandleFunc("/signin", Signin)
	//http.HandleFunc("/welcome", Welcome)
	//http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	//log.Fatal(http.ListenAndServe(":8000", nil))

	r := gin.Default()
	//testing ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signin", Signin)
	r.GET("/welcome", Welcome)

	r.Run() // listen and serve on 0.0.0.0:8080
}
