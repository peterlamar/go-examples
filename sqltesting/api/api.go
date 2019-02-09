package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peterlamar/go-examples/sqltesting/pythonmovies"
	log "github.com/sirupsen/logrus"
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
	router.GET("/getboxdifference/:arg", GetBoxDifference)

	return router
}

// GetBoxDifference Get Box office difference, more of an example
// to illustrate sql testing
func GetBoxDifference(ctx *gin.Context) {

	inputArg, err := strconv.Atoi(ctx.Param("arg"))

	if err != nil {
		log.Fatal(err)
	}

	difference := GetDifference(inputArg)

	ctx.JSON(http.StatusOK, gin.H{
		"Hello": difference,
	})
}

// GetDifference function to be tested
func GetDifference(inputArg int) (rtn int) {

	log.Info("input ", inputArg)

	movieInfo := pythonmovie.GetMovieInfo(inputArg)

	rtn = movieInfo.WordwideBox - movieInfo.DomesticBox

	log.Info("Difference is ", rtn)

	return
}
