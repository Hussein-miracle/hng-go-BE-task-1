package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type SubmissionDetails struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	GithubFileUrl string `json:"github_file_url"`
	Track         string `json:"track"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func handleGetSubmissionDetails(c *gin.Context) {
	slack_name := c.Query("slack_name")
	track := c.Query("track")

	utc_time := time.Now().UTC()

	details := SubmissionDetails{
		SlackName:     slack_name,
		Track:         track,
		CurrentDay:    utc_time.Weekday().String(),
		UtcTime:       utc_time.Format("2023-09-02T15:04:05Z"),
		GithubFileUrl: "https://github.com/Hussein-miracle/hng-go-BE-task-1/blob/master/main.go",
		GithubRepoUrl: "https://github.com/Hussein-miracle/hng-go-BE-task-1",
		StatusCode:    200,
	}

	// fmt.Println("%+v", details)

	c.JSON(http.StatusOK, details)
}

func main() {
	router := gin.Default()
	// fmt.Println("%v", router)
	// curr := utc.Now()

	// fmt.Println(time.Now().UTC().Format("2023-09-02T15:04:05Z"))
	// fmt.Println(time.Now().Weekday().String())

	router.GET("/api", handleGetSubmissionDetails)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := router.Run(":" + port)

	if err != nil {
		log.Panicf("Error Ocurred: %s", err)
	}

}
