package main

import (

	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


func main() {

	router := gin.Default()

	go router.GET("/albums", getAlbums)
    go router.GET("/albums/:id", getAlbumByID)
	go router.POST("/albums", postAlbums)
	go router.DELETE("/albums/:id",deleteAlbumbyID)

	router.Run("localhost:7000")
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 69.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 19.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumID(c *gin.Context){

	id := c.Param("id")

	for _,a := range albums {
		if a.ID == id {
			fmt.Println(a.ID)
		}
	}
}

func deleteAlbumbyID(c *gin.Context){

	if := c.Param("id")
	fmt.Println("Enter position to delete::")
	fmt.Scanln(&pos)

	new_arr := make([]album, (len(arr) - 1))
	k := 0
	for i:= 0; i< (len(arr) - 1); {
		if i != pos {
			new_arr[i] = album[k]
			k++
			i++
		} else {
			k++
		}
	}
	album[] = new_arr[]

	for i:= 0; i < (len(album) - 1); i++ {

		fmt.Println(new_Arr[i])
	}
}

