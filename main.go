package main

import (
	"github.com/aldiui/golang-mrt/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	initiateRouter()
}

func initiateRouter() {
	var router = gin.Default()
	var api = router.Group("/v1/api")

	station.Initiate(api)
	router.Run(":1020")
}
