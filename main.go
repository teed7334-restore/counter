package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"./services"

	"./beans"
	"./env"
	"./route"
	"github.com/bitly/go-nsq"
)

var cfg = env.GetEnv()

func main() {
	services.Create("INFO", "存入Elasticsearch")
	upStream := make(chan time.Time)
	go func() {
		webService()
		upStream <- time.Now()
	}()
	go func() {
		messageService()
		upStream <- time.Now()
	}()
	<-upStream
}

//webService Restful API服務
func webService() {
	api := route.API()
	api.Run(":8805")
}

//messageService Message Quete服務
func messageService() {
	host := cfg.Message.Received.Address
	InitConsumer(cfg.Message.Topic, cfg.Message.Channel, host)
}

//InitConsumer 初始化消費者
func InitConsumer(topic string, channel string, host string) bool {
	interval := time.Second * 2
	upStream := make(chan int, 1)
	config := nsq.NewConfig()
	config.LookupdPollInterval = interval
	query, err := nsq.NewConsumer("Mail", "SendMail", config)
	if err != nil {
		log.Panic(err)
	}
	query.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		chat := string(message.Body)
		quete := getQuete(chat)
		runServices(quete)
		return nil
	}))
	if err = query.ConnectToNSQLookupd(host); err != nil {
		panic(err)
	}
	<-upStream
	return true
}

func getQuete(message string) []string {
	quete := strings.Split(message, "</UseService>")
	return quete
}

func runServices(quete []string) {
	useService := quete[0]
	params := []byte(quete[1])
	switch useService {
	case "SendMail":
		sendMail := &beans.SendMail{}
		json.Unmarshal(params, sendMail)
		services.SendMail(sendMail)
	}
}
