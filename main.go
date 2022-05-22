package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CounterMiddleware(counters *Counters, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		counters.Incr(key, 1)
		c.Next()
	}
}

func Flush2Stderr(counters *Counters, key string) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		fmt.Println(counters.ResetByKey(key))
	}
}

func main() {
	counters := NewCounters()
	counters.Init("ping", Flush2Stderr)
	r := gin.Default()
	r.GET("/ping", CounterMiddleware(counters, "ping"), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
