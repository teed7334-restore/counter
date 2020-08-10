package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/teed7334-restore/counter/env"
	"github.com/teed7334-restore/counter/route"

	_ "github.com/joho/godotenv/autoload"

	nsq "github.com/nsqio/go-nsq"
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
	host := os.Getenv("message.received.address")
	topic := os.Getenv("message.topic")
	channel := os.Getenv("message.channel")
	InitConsumer(topic, channel, host)
}

//InitConsumer 初始化消費者
func InitConsumer(topic string, channel string, host string) bool {
	upStream := make(chan int, 1)
	for _, worker := range cfg.Works.Worker {
		go func(worker *env.Worker) {
			mQuery(host, worker.Topic, worker.Channel, worker.Interval)
		}(worker)
	}
	<-upStream
	return true
}

//mQuery 啟用單一消費者
func mQuery(host string, topic string, channel string, interval time.Duration) *nsq.Consumer {
	config := nsq.NewConfig()
	config.LookupdPollInterval = interval
	query, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Println(err)
	}
	query.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		chat := string(message.Body)
		quete := getQuete(chat)
		runServices(quete)
		return nil
	}))
	err = query.ConnectToNSQLookupd(host)
	if err != nil {
		log.Println(err)
	}
	return query
}

//getQuete 取得工作內容
func getQuete(message string) []string {
	quete := strings.Split(message, "</UseService>")
	return quete
}

//runServices 運行系統服務
func runServices(quete []string) {
	path := quete[0]
	params := []byte(quete[1])
	host := os.Getenv("housekeeper.host")
	url := host + "/" + path
	postURL(url, params)
}

//postURL 將資料打到REST API
func postURL(url string, params []byte) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	request.Header.Set("X-Custom-Header", "counter")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
