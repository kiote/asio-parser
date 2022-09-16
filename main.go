package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "github.com/gin-gonic/gin"
)

type oneLecture struct {
    Date string `json:"date"`
    Time string `json:"time"`
    Room string `json:"room"`
    Code string `json:"code"`
}


// http://www.tlu.ee/masio/?id=aine&aine=IFI8110.DT&time=1663794000&keel=1
func getLecturesByCode(code string)([]oneLecture) {
    return []oneLecture{
       {Date: "Wednesday 22. September", Time: "10:00-13:00", Room: "A-343", Code: code},
    }
}

// getLectures responds with the list of all lectures this week as JSON.
func getLectures(c *gin.Context) {
    code := c.Param("code")
    lectures := getLecturesByCode(code)
    c.IndentedJSON(http.StatusOK, lectures)
}

func pageContent(link string)(string) {
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
    router.GET("/lectures/:code", getLectures)

    router.Run("localhost:8080")
}
