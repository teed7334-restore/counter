package beans

//Response MQ回應參數
type Response struct {
	Status  bool   //`json:"Status"`
	Channel string //`json:"Channel"`
	Message string //`json:"Message"`
}
