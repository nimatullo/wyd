package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"wyd/activity"
	"wyd/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	activity.CURRENT_ACTIVITY = database.GetCurrentActivityFromDb()
	fmt.Println("Current activity:", activity.CURRENT_ACTIVITY)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", HelloWorld)
	router.POST("/activity", UpdateCurrentActivity)
	router.GET("/activity", GetCurrentActivity)
	router.GET("/stream", StreamHandler)

	port := os.Getenv("PORT")

	if len(port) > 0 {
		fmt.Println("Running server on port", port)
		router.Run(":" + port)
	} else {
		fmt.Println("Running server on port default port")
		router.Run(":8080")
	}
}

func HelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func UpdateCurrentActivity(c *gin.Context) {
	updateJson := activity.Activity{}

	c.BindJSON(&updateJson)
	updateJson.Since = time.Now().Format("2006-01-02 15:04:05")

	if database.UpdateCurrentActivityInDb(updateJson) {
		updateJson.Ready = true
		activity.CURRENT_ACTIVITY = updateJson
	}

	c.IndentedJSON(http.StatusOK, activity.CURRENT_ACTIVITY)
}

func GetCurrentActivity(c *gin.Context) {
	currentActivity := database.GetCurrentActivityFromDb()

	c.IndentedJSON(http.StatusOK, currentActivity)
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	
	// When a connection is first made to the event-stream, we need to publish the last known activity.
	// this boolean lets the stream function to know whether or not the initial activity should be published.
	firstStream := true

	c.Stream(func(w io.Writer) bool {
		if firstStream && len(activity.CURRENT_ACTIVITY.Name) > 0 {
			firstStream = false
			Send(w)
		}

		if activity.CURRENT_ACTIVITY.Ready {
			Send(w)
		}
		return true
	})

	fmt.Println("Closing connection")
}

func Send(w io.Writer) {
	bytes, err := json.Marshal(activity.CURRENT_ACTIVITY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "data: %s\n\n", string(bytes))
	activity.CURRENT_ACTIVITY.Ready = false
}
