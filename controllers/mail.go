package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../beans"
	"../env"
	"github.com/bitly/go-nsq"
	"github.com/gin-gonic/gin"
)

var cfg = env.GetEnv()

func postMessage(message string) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(cfg.Message.Post.Address, config)
	err := w.Publish("Mail", []byte(message))
	if err != nil {
		fmt.Println(err)
	}
	w.Stop()
}

//SendMail 寄信用API
func SendMail(c *gin.Context) {
	sendMail := &beans.SendMail{}
	err := c.BindJSON(sendMail)
	if err != nil {
		log.Println(err)
	}
	json, _ := json.Marshal(sendMail)
	quete := "SendMail</UseService>" + string(json)
	postMessage(quete)
	c.JSON(http.StatusOK, gin.H{"status": "true"})
}
