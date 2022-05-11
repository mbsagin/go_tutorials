package main

import (
	"encoding/json"
	"errors"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	router.GET("/album/:id", getAlbumById)
	router.GET("/albums", getAllAlbums)
	
	router.POST("/album", postAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbumById(c *gin.Context) {
	// Requested album id
	albumId := c.Param("id")

	for _, a := range albums {
		if a.ID == albumId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	errorMessage := "Album could not found with the ID " + albumId;
	logError(errors.New(errorMessage))

	c.IndentedJSON(http.StatusNotFound, gin.H{ "isError": true, "message": errorMessage })
}

// GET all the albums
func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST an album
func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)

    if err != nil {
		logError(err)
		
		c.IndentedJSON(http.StatusUnprocessableEntity, newAlbum)

		return
    }

    albums = append(albums, newAlbum)

    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// POST albums
func postAlbums(c *gin.Context) {
	var newAlbums []album

	err := c.BindJSON(&newAlbums)

    if err != nil {
        logError(err)
		
		c.IndentedJSON(http.StatusUnprocessableEntity, newAlbums)

		return
    }

    albums = append(albums, newAlbums...)

    c.IndentedJSON(http.StatusCreated, newAlbums)
}

// JSON file reader
func readJson(fileName string) []byte {
	jsonFile, err := os.Open(fileName)

	fmt.Println(jsonFile)

	if err != nil {
		logError(err)
	}
	
	bytes, err := ioutil.ReadAll(jsonFile)
	
	jsonFile.Close()

	if err != nil {
		logError(err)
	}

	return bytes
}

// Error logger, TODO: Log errors to DB with traceId
func logError(err error) {
	traceId := uuid.New()

	log.SetPrefix("["+traceId.String()+"]")
    log.SetFlags(0)
	

	log.Println(err)

	// TODO: DB insert
}
