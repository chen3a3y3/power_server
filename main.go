package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"private/power_server/service"
)

func main() {
	r := gin.Default()
	r.GET("ping", service.Ping)

	if err := r.Run("localhost:8080"); err != nil {
		fmt.Printf("failed to start gin server, err = %v\n", err)
	}
}
