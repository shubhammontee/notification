package app

import (
	"album-manager/notification/handler"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	handler := handler.NewHandler()
	router.POST("/api/v1/data", handler.PostDataToKafka)
	router.Run(":9000")

}
