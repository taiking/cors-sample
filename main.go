package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.POST("/hoge", func(c *gin.Context) {
    	c.Writer.Header().Set("X-Hoge", "hogehoge")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.String(200, "{}")
    })
    r.OPTIONS("/hoge", func(c *gin.Context) {
    	 c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    	 c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
    	 c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
    	 c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
  		c.Writer.WriteHeader(200)
    	 c.Abort()
    })
    r.Run(":3001")
}