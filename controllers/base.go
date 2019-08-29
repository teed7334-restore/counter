package controllers

import (
	"fmt"
	"log"

	"github.com/bitly/go-nsq"
	"github.com/teed7334-restore/counter/env"

	"github.com/gin-gonic/gin"
)

var cfg = env.GetEnv()

//將要執行的排程寫到MQ postMessage
func postMessage(channel string, message string) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(cfg.Message.Post.Address, config)
	err := w.Publish(channel, []byte(message))
	if err != nil {
		fmt.Println(err)
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
