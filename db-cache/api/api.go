package api

import (
	"github.com/gin-gonic/gin"
	"github.com/peterlamar/go-examples/db-cache/pythonmovies"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// Helloarg is used for injecting example post data
type Helloarg struct {
	Helloargstr string `json:"helloarg"`
}

// Heartbeat A wait to determine if the service is still alive
func Heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Still": "Alive",
	})
}

// SetupRouter is used to setup the router paths
func SetupRouter() *gin.Engine {
	router := gin.Default() // Create router
	router.GET("/", Heartbeat)
	router.GET("/helloget/:arg", Helloget)

	return router
}

// Helloget This is a get example
func Helloget(ctx *gin.Context) {

	inputArg, err := strconv.Atoi(ctx.Param("arg"))

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	// Get the movie name, the first call will be a cache miss and hit the db
	name := pythonmovie.GetMovieName(inputArg)

	log.Printf("GetMovieName DB took %s", time.Since(start))

	start2 := time.Now()

	// Get the movie name, the second call will be a cache hit
	name = pythonmovie.GetMovieName(inputArg)

	log.Printf("GetMovieName Cache took %s", time.Since(start2))

	ctx.JSON(http.StatusOK, gin.H{
		"Hello": name,
	})
}
