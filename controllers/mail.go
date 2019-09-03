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
	response := doSendMail(params)
	c.JSON(http.StatusOK, response)
}

//doSendMail 運行寄信用API
func doSendMail(params *beans.SendMail) *beans.Response {
	jsonStr, _ := json.Marshal(params)
	channel := "Mail"
	message := "Mail/SendMail</UseService>" + string(jsonStr)
	postMessage(channel, message)
	response := &beans.Response{Status: true, Channel: channel, Message: message}
	return response
}
