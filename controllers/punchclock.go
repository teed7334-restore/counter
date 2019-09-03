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
	response := doUploadDailyPunchclockData(params)
	c.JSON(http.StatusOK, response)
}

//doUploadDailyPunchclockData 運行將每天員工打卡資料上鏈
func doUploadDailyPunchclockData(params *homekeeper.DailyPunchclockData) *beans.Response {
	jsonStr, _ := json.Marshal(params)
	message := "PunchClock/UploadDailyPunchclockData</UseService>" + string(jsonStr)
	channel := "PunchClock"
	postMessage(channel, message)
	response := &beans.Response{Status: true, Channel: channel, Message: message}
	return response
}
