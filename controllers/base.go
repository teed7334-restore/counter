package controllers

import (
	"log"
	"os"

	nsq "github.com/nsqio/go-nsq"

	"github.com/gin-gonic/gin"
)

//將要執行的排程寫到MQ postMessage
func postMessage(channel string, message string) {
	config := nsq.NewConfig()
	address := os.Getenv("message.post.address")
	w, _ := nsq.NewProducer(address, config)
	err := w.Publish(channel, []byte(message))
	if err != nil {
		log.Println(err)
	}
	w.Stop()
}

//getParams 取得HTTP POST帶過來之參數
func getParams(c *gin.Context, params interface{}) {
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
}
