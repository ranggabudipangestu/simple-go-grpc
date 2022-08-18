package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggabudipangestu/go-grpc-client/models"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	returnedData := models.ResponseWrapper{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       data,
	}

	c.JSON(http.StatusOK, returnedData)
}

func HandleError(c *gin.Context, status int, message string) {
	returnedData := models.ResponseWrapper{
		Message:    message,
		StatusCode: status,
		Success:    false,
	}

	c.JSON(status, returnedData)
}

func HandleSuccessReturn(data interface{}) string {
	var returnData = models.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}
	res, _ := json.Marshal(&returnData)
	return string(res)
}
