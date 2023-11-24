package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestToBody(r *http.Request, result any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err, "Error Decode on helper at ReadRequestToBody")
}
func WriteRequestToBody(w http.ResponseWriter, response any) {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err, "Error Encode on helper at WriteRequestToBody")
}
