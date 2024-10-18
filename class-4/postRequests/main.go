package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type RequestFormat struct {
	UnixTime int64 `json:"time"`
}

func main() {
	r := gin.Default()
	r.POST("/convert/time", handleConversion)
	r.Run()
}

func handleConversion(c *gin.Context) {
	var request RequestFormat
	c.BindJSON(&request)
	// request.Time is time in epoch format
	requestTime := time.Unix(request.UnixTime, 0)
	fmt.Println(requestTime)
	// convert to a user readable format and return
	humanReadableTime := requestTime.Format("2006-01-02 15:04:05")
	fmt.Println(humanReadableTime)
	c.JSON(200, gin.H{"time": humanReadableTime})
}
