package response

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

func WriteJson(writer http.ResponseWriter, data interface{}) {
	jsonResponse, _ := json.Marshal(data)
	fmt.Fprintf(writer, "%s", jsonResponse)
}

func DecodeJson(r io.Reader, dst interface{}) {
	d := json.NewDecoder(r)
	d.Decode(dst)
}

func DecodeJsonString(s string, dst interface{}) {
	json.Unmarshal([]byte(s), dst)
}