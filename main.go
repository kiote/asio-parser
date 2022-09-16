package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "log"
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

// http://www.tlu.ee/masio/?id=aine&aine=IFI8110.DT&time=1663794000&keel=1

// getLectures responds with the list of all lectures this week as JSON.
func getLectures(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, lectures)
}

func PageContent(link string)(string) {
    res, err := http.Get(link)
    if err != nil {
        log.Fatal(err)
    }
    content, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func main() {
    router := gin.Default()
    router.GET("/lectures", getLectures)

    router.Run("localhost:8080")
}
