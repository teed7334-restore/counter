# counter
佇列管理員，可以運用在做系統非同步、資料庫延遲寫入、寄信、區塊鏈上鏈等功能上

## 資料夾結構
beans 用來裝Call API後之ResultObject

controllers Restful API呼叫用控制器

dev_env Docker-Compose環境產生設定

.env 系統設定

route 系統路由設定

main.go 主程式

## 程式運行原理
本系統會透過多行程啟動一個Restful API與一個與Messae Quete溝通，並查看任務且執行任務之服務

當使用者透過呼叫本系統Restful API時，系統會將需要操作的事務記錄到Message Quete去，當使用者收到HTTP 200之後，本系統會每隔二秒自己去查看需要執行的任務，並執行他

本系統透過nsq做為Message Quete

本系統可以搭配服務管理員一起使用

https://github.com/teed7334-restore/homekeeper

如果連nsq都不想自己架的人，可以自己安裝Docker與Docker Compose，自己到./dev_env資料夾下打docker-compose up -d --build，nsq會自己架好

## 程式操作流程
1. 將.env.swp檔名改成.env
2. 修改.env並設定您的Message Quete Server
3. go run main.go

## API呼叫網址與參數
寄信服務(須搭配服務管理員) http://[Your Host Name]:8805/Mail/SendMail
```
//HTTP Header需設定成Content-Type: application/json
{
    "to": "admin@admin.com",
    "subject": "這是一封測試信",
    "content": "這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />"
}
```