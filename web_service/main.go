package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album {}

func init() {
	byteAlbumData := readJson("data.json")

	json.Unmarshal(byteAlbumData, &albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)

    if err != nil {
        fmt.Println(err)
		return
    }

    albums = append(albums, newAlbum)

    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func readJson(fileName string) []byte {
	jsonFile, err := os.Open(fileName)

	fmt.Println(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	
	bytes, err := ioutil.ReadAll(jsonFile)
	
	jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

