package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/teed7334-restore/counter/route"

	"github.com/teed7334-restore/counter/env"

	"github.com/bitly/go-nsq"
)

var cfg = env.GetEnv()

func main() {
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
	upStream := make(chan int, 1)
	go func() {
		interval := time.Second * 10
		mQuery(host, "PunchClock", "UploadDailyPunchclockData", interval)
	}()
	go func() {
		interval := time.Second * 2
		mQuery(host, "Mail", "SendMail", interval)
	}()
	<-upStream
	return true
}

func mQuery(host string, topic string, channel string, interval time.Duration) *nsq.Consumer {
	config := nsq.NewConfig()
	config.LookupdPollInterval = interval
	query, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Panic(err)
	}
	query.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		chat := string(message.Body)
		quete := getQuete(chat)
		runServices(quete)
		return nil
	}))
	err = query.ConnectToNSQLookupd(host)
	if err != nil {
		panic(err)
	}
	return query
}

func getQuete(message string) []string {
	quete := strings.Split(message, "</UseService>")
	return quete
}

func runServices(quete []string) {
	path := quete[0]
	params := []byte(quete[1])
	url := cfg.Housekeeper.Host + "/" + path
	postURL(url, params)
}

func postURL(url string, params []byte) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	request.Header.Set("X-Custom-Header", "counter")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
}
