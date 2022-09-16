package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "github.com/gin-gonic/gin"
    "strconv"
    "time"
)

type oneLecture struct {
    Date string `json:"date"`
    Time string `json:"time"`
    Room string `json:"room"`
    Code string `json:"code"`
}

var uriPartEng = "keel=1"
var link = "http://www.tlu.ee/masio/?id=aine&aine="
var nextWeekSec = int64(800608)

func uriPartCurrentTime() (string) {
    t := time.Now()
    sec := t.Unix() + nextWeekSec
    return "time=" + strconv.Itoa(int(sec))
}

func parseLectureInfo(pageContent string)(oneLecture) {
    lecture := oneLecture{Date: "some date", Time: "some time", Room: "some room", Code: "some code"}
    return lecture 
}

func getLectureByCode(code string)([]oneLecture) {
    url := link + code + "&" + uriPartCurrentTime() + "&" + uriPartEng
    log.Print("Requesting: " + url)
    pageContent := getPageContent(url)
    lecture := parseLectureInfo(pageContent)
    log.Print(pageContent)
    return []oneLecture{
       lecture,
    }
}

// getLectures responds with the list of all lectures this week as JSON.
func getLectures(c *gin.Context) {
    code := c.Param("code")
    lectures := getLectureByCode(code)
    c.IndentedJSON(http.StatusOK, lectures)
}

func getPageContent(link string)(string) {
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
