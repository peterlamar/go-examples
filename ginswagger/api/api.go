package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	"net/http"
)

// Helloarg is used for injecting example post data
type Helloarg struct {
	Helloargstr string `json:"helloarg"`
}

// Heartbeat godoc
// @Summary A wait to determine if the service is still alive
// @Produce json
// @Success 200 {object} object
func Heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Still": "Alive",
	})
}

// SetupRouter is used to setup the router paths
func SetupRouter() *gin.Engine {
	router := gin.Default() // Create router
	router.GET("/", Heartbeat)
	router.POST("/hellopost", Hellopost)
	router.GET("/helloget/:arg", Helloget)
	// To disable if an env variable is set.
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER_MODE"))
	return router
}

// Hellopost This is a post example
// @Summary This is used to show a post example
// @Accept json
// @Produce json
// @Param Context body Helloarg true "A json object with hello"
// @Success 200 {integer} integer
// @Router /hellopost [post]
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

// Helloget This is a get example
// @Summary This is used to show a get example
// @Accept json
// @Produce json
// @Param arg path integer true "Your favorite integer"
// @Success 200 {integer} integer
// @Router /helloget/{arg} [get]
func Helloget(ctx *gin.Context) {
	// Unmarshal the json request to a struct
	int := ctx.Param("arg")

	fmt.Println("my arg ", int)

	ctx.JSON(http.StatusOK, gin.H{
		"Hello": int,
	})
}
