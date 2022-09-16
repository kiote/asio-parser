package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type oneLecture struct {
	Date string `json:"date"`
	Time string `json:"time"`
	Room string `json:"room"`
}

var lectures = []oneLecture{
	{Date: "Wednesday 22. September", Time: "10:00-13:00", Room: "A-343"},
}

// getAlbums responds with the list of all albums as JSON.
func getLectures(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, lectures)
}

func main() {
    router := gin.Default()
    router.GET("/lectures", getLectures)

    router.Run("localhost:8080")
}
