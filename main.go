package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"time"

	"fmt"

	"os"
)

var CURRENT_ACTIVITY Activity = Activity{
	Name:    "",
	Website: "",
	Since:   "",
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", helloWorld)
	router.POST("/activity", updateCurrentActivity)
	router.GET("/activity", getCurrentActivity)

	port := os.Getenv("PORT")

	router.Run(":" + port)
}

type Activity struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Since   string `json:"since"`
}

func helloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func updateCurrentActivity(c *gin.Context) {
	var activity Activity
	// Parse the JSON into the struct and add current time
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activity.Since = time.Now().Format("2006-01-02 15:04:05")

	c.BindJSON(&activity)
	CURRENT_ACTIVITY = activity

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Activity updated",
		"current": CURRENT_ACTIVITY,
	})

	fmt.Println(CURRENT_ACTIVITY)
}

func getCurrentActivity(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"current": CURRENT_ACTIVITY,
	})
}
