package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pdata struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

var pdatas = []pdata{
	{Name: "maruto", Pass: "maruto"},
}

func main() {
	engine := gin.Default()
	engine.GET("/", helloTest)
	engine.GET("/show", getAllpdata)
	engine.POST("/post", postPdatas)
	engine.Run(":3000")
}

func helloTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
func getAllpdata(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, pdatas)
}
func postPdatas(c *gin.Context) {
	var newPdata pdata

	if err := c.BindJSON(&newPdata); err != nil {
		return
	}

	// Add the new album to the slice.
	pdatas = append(pdatas, newPdata)
	for _, i := range pdatas {
		log.Println(i)
	}
	c.IndentedJSON(http.StatusCreated, newPdata)
}
