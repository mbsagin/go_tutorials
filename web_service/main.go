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

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	byteAlbumData := readJson("data.json")

	var albums []album

	json.Unmarshal(byteAlbumData, &albums)

	c.IndentedJSON(http.StatusOK, albums)
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

