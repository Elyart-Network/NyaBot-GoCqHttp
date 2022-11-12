package ApiFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetRequest Send Get Requests to GoCqHttp
func getRequest(Endpoint string) {

}

// PostRequest Send Post Requests to GoCqHttp
func postRequest(Endpoint string, Params interface{}) (Context []byte) {
	byteSlice, _ := json.MarshalIndent(Params, "", "")
	Request, err := http.Post(data.CqHttpHost+Endpoint, "application/json", bytes.NewBuffer(byteSlice))
	if err != nil {
		log.Println(err)
	}
	defer Request.Body.Close()
	Context, _ = io.ReadAll(Request.Body)
	return Context
}
