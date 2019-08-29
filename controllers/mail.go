package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/teed7334-restore/counter/beans"

	"github.com/gin-gonic/gin"
)

//SendMail 寄信用API
func SendMail(c *gin.Context) {
	params := &beans.SendMail{}
	getParams(c, params)
	json, _ := json.Marshal(params)
	channel := "Mail"
	message := "Mail/SendMail</UseService>" + string(json)
	postMessage(channel, message)
	response := &beans.Response{Status: true, Channel: channel, Message: message}
	c.JSON(http.StatusOK, response)
}
