package main

import (

	"net/http"

	"github.com/gin-gonic/gin"

	"web-service-gin/Albums"

	"fmt"

	"sync"
)


// album represents data about a record album.
type album struct {

//    mutex sync.Mutex
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 69.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 19.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

//    c.mutex.Lock()
    c.IndentedJSON(http.StatusOK, albums)
//    defer c.mutex.Unlock()
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {

//    c.mutex.Lock()
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
//    defer c.mutex.Unlock()
}

func getAlbumByID(c *gin.Context) {

//    c.mutex.Lock()
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
//    c.mutex.Unlock()
}

func updateAlbumID(c *gin.Context){

//	c.mutex.Lock()
	id := c.Param("id")

	for _,a := range albums {
		if a.ID == id {
			fmt.Println(a.ID)
		}
	}
//        defer c.mutex.Unlock()
}

func deleteAlbumbyID(c *gin.Context){

//        c.mutex.Lock()
	id := c.Param("id")

	var pos int = 0
	fmt.Println("Enter position to delete::")
	fmt.Scanln(&pos)

	new_arr := make([]album, (len(albums) - 1))

	k := 0
	for _,a := range albums {
		if a.ID == id{
			for i:= 0; i< (len(albums) - 1); {
				if i != pos {
					new_arr[i] = albums[k]
					k++
					i++
				} else {
				k++
				}
			}
//	albums[pos] := new_arr[pos]

	for i:= 0; i < (len(albums) - 1); i++ {

		fmt.Println(new_arr[i])
	}
//        defer c.mutex.Unlock()
}

func() {

	router := gin.Default()

	var wg sync.WaitGroup

	wg.Add(3)

	func(){

		go router.GET("/albums", getAlbums)
        	go router.GET("/albums/:id", getAlbumByID)
		go router.POST("/albums", postAlbums)
        }()

	wg.Wait()
	router.Run("localhost:6000")
}
