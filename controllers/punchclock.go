package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teed7334-restore/counter/beans"
	homekeeper "github.com/teed7334-restore/homekeeper/beans"
)

//UploadDailyPunchclockData 將每天員工打卡資料上鏈
func UploadDailyPunchclockData(c *gin.Context) {
	params := &homekeeper.DailyPunchclockData{}
	getParams(c, params)
	json, _ := json.Marshal(params)
	message := "PunchClock/UploadDailyPunchclockData</UseService>" + string(json)
	channel := "PunchClock"
	postMessage(channel, message)
	response := &beans.Response{Status: true, Channel: channel, Message: message}
	c.JSON(http.StatusOK, response)
}
