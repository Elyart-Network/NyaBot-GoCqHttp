package gcqapi

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"bytes"
	"io"
	"log"
	"net/http"
)

type respStruct struct {
	Status  string      `json:"status"`
	RetCode int         `json:"retcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Echo    string      `json:"echo"`
}

// GetRequest Send Get Requests to GoCqHttp
func getRequest(Endpoint string) (Context []byte) {
	Request, err := http.Get(gcqdata.CqHttpHost + Endpoint)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(Request.Body)
	Context, _ = io.ReadAll(Request.Body)
	return Context
}

// PostRequest Send Post Requests to GoCqHttp
func postRequest(Endpoint string, Params interface{}) (Context []byte) {
	byteSlice, _ := gcqdata.JsonLib.MarshalIndent(Params, "", "")
	Request, err := http.Post(gcqdata.CqHttpHost+Endpoint, "application/json", bytes.NewBuffer(byteSlice))
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(Request.Body)
	Context, _ = io.ReadAll(Request.Body)
	return Context
}

// RespDecoder Decode Json respond to Only Data
func respDecoder(Context []byte) []byte {
	RespStruct := respStruct{}
	err := gcqdata.JsonLib.Unmarshal(Context, &RespStruct)
	if err != nil {
		log.Println(err)
	}
	byteSlice, err := gcqdata.JsonLib.Marshal(RespStruct.Data)
	if err != nil {
		log.Println(err)
	}
	return byteSlice
}
