package handler

import (
	"album-manager/notification/kafka"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AlbumHandlerInterface ...
type AlbumHandlerInterface interface {
	PostDataToKafka(c *gin.Context)
	// CreateAlbum(c *gin.Context)
	// DeleteAlbum(c *gin.Context)
	// DeleteImage(c *gin.Context)
}
type albumHandler struct {
	//service service.Service
}

//NewHandler ...
func NewHandler() AlbumHandlerInterface {
	return &albumHandler{
		//service: service,
	}
}

func (album albumHandler) PostDataToKafka(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	form := &struct {
		Message string `form:"message" json:"message"`
	}{}
	if err := ctx.ShouldBindJSON(form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while marshalling json: %s", err.Error()),
			},
		})

	}
	//ctx.Bind(form)
	formInBytes, err := json.Marshal(form)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while marshalling json: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	err = kafka.Push(parent, nil, formInBytes)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while push message into kafka: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success push data into kafka",
		"data":    form,
	})
}
