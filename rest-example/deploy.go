package main

// Experimental vendoring used for dependencies
// https://github.com/golang/go/commit/183cc0cd41f06f83cb7a2490a499e3f9101befff
import (
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
)

// Main entrypoint of REST api

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/v1/deploy").To(deploy))
	ws.Route(ws.GET("/v1/verify").To(verify))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

// Example Curl for testing
// curl -H "Content-Type: application/xml" -X GET http://localhost:8080/v1/deploy


// GET http://localhost:8080/v1/deploy
func deploy(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "return list of deployments")
	// Hook for deploy
}

// GET http://localhost:8080/v1/verify
func verify(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "verify target environment")
	// Hook for verify
}