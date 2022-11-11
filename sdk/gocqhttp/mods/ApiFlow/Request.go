package ApiFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/GcqServer"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetRequest Send Get Requests to GoCqHttp
func GetRequest(Endpoint string) {

}

// PostRequest Send Post Requests to GoCqHttp
func PostRequest(Endpoint string, Params interface{}) (Context []byte) {
	byteSlice, _ := json.MarshalIndent(Params, "", "")
	Request, err := http.Post(GcqServer.CqHttpHost+Endpoint, "application/json", bytes.NewBuffer(byteSlice))
	if err != nil {
		log.Println(err)
	}
	defer Request.Body.Close()
	Context, _ = io.ReadAll(Request.Body)
	return Context
}
