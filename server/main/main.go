package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("a", "127.0.0.1", "IP Address")

	engine := gin.Default()
	engine.GET("/action", func(c *gin.Context) {
		// 制御を実行
		action()
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	})

	engine.Run(*addr + ":55555")
}

func action() {

}
