package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func ResponseJson(w http.ResponseWriter, p interface{}) {
	ubahkeByte, err := json.Marshal(p)
	if err!=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(`content-type`,`application/json`)
	w.WriteHeader(200)
	w.Write(ubahkeByte)
}
func IsError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return true
	}
	return false
}
