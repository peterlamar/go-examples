package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	router.POST("/hellopost", Hellopost)

	return router
}

// Helloget This is a get example
func Helloget(ctx *gin.Context) {
	// Unmarshal the json request to a struct
	int := ctx.Param("arg")

	log.Info("my arg ", int)

	ctx.JSON(http.StatusOK, gin.H{
		"Hello": int,
	})
}

// Hellopost This is a post example
func Hellopost(ctx *gin.Context) {
	// Unmarshal the json request to a struct
	var json Helloarg

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Hello": "world",
	})
}
