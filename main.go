package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var CURRENT_ACTIVITY Activity = Activity{
	Name:    "",
	Website: "",
	Since:   "",
	Ready:   false,
}

func main() {
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

type Activity struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Since   string `json:"since"`
	Ready   bool   `json:"ready"`
}

func HelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func UpdateCurrentActivity(c *gin.Context) {
	var activity Activity
	activity.Since = time.Now().Format("2006-01-02 15:04:05")

	c.BindJSON(&activity)
	CURRENT_ACTIVITY = activity
	CURRENT_ACTIVITY.Ready = true

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Activity updated",
		"current": CURRENT_ACTIVITY,
	})

	fmt.Println(CURRENT_ACTIVITY)
}

func GetCurrentActivity(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"current": CURRENT_ACTIVITY,
	})
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")

	c.Stream(func(w io.Writer) bool {
		if CURRENT_ACTIVITY.Ready {
			bytes, err := json.Marshal(CURRENT_ACTIVITY)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(w, "data: %s\n\n", string(bytes))
			CURRENT_ACTIVITY.Ready = false
		}
		return true
	})

	fmt.Println("Closing connection")
}
