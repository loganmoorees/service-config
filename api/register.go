package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"service-config/core"
)

func register(c *gin.Context) {
	log.Println("Register Initializing Executor")
	var request core.Instance
	if e := c.ShouldBindJSON(&request); e != nil {
		log.Println("Register Parameter Processing Error: ", e)
	}
	// key is appid + env
	initRegistry().Register(request, request.LatestTimestamp)

}

func initRegistry() *core.Registry {
	return &core.Registry{
		Apps: make(map[string]*core.Application),
	}
}

func InitApplication(appId string) *core.Application {
	return &core.Application{
		Appid:    appId,
		Instance: make(map[string]*core.Instance),
	}
}
