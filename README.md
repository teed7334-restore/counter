# counter
佇列管理員

## 資料夾結構
beans 用來裝Call API後之ResultObject

controllers Restful API呼叫用控制器

dev_env Docker-Compose環境產生設定

env 系統設定

route 系統路由設定

services 系統任務執行主程式

main.go 主程式

## 程式運行原理
本系統會透過多行程啟動一個Restful API與一個與Messae Quete溝通，並查看任務且執行任務之服務

當使用者透過呼叫本系統Restful API時，系統會將需要操作的事務記錄到Message Quete去，當使用者收到HTTP 200之後，本系統會每隔二秒自己去查看需要執行的任務，並執行他

本系統透過nsq做為Message Quete

目前它己有一個現成的寄信服務，其他可以依自己做擴充，要修改Resful API，就自己在Route/Web.go裡面設定路由，然後將對應的程式放到Controllers去，要開發供MQ執行的任務，就將程式放到Services去，並修改./main.go的runService內容

本來有想過這隻程式只做佇列排程就好，最後在透過gRPC或是CURL去呼叫其他的Restful API，這樣會比較乾淨，只是透過外部呼叫效能不會比同一個系統裡面執行還好

如果連nsq都不想自己架的人，可以自己安裝Docker與Docker Compose，自己到./dev_env資料夾下打docker-compose up -d --build，nsq會自己架好

## 必須套件
本程式透過Google Protobuf 3產生所需之ResultObject，然Proto 3之後官方不支持Custom Tags，所以還需要多安裝一個寫入retags的套件

git clone https://github.com/qianlnk/protobuf.git $GOPATH/src/github.com/golang/protobuf

go install $GOPATH/src/github.com/golang/protobuf/protoc-gen-go

還有與nsq溝通用之套件

go get -u -v github.com/bitly/go-nsq

及Restful Framework

go get -u -v github.com/gin-gonic/gin

## 程式操作流程
1. 將./env/env.swp檔名改成env.go
2. 修改./env/env.go並設定您的SMTP Server與Message Quete Server
3. 到./beans底下，運行protoc --go_out=plugins=grpc+retag:. *.proto
4. go run main.go