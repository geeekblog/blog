package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	err := r.Run(":8081")
	if err != nil {
		return
	}
}
