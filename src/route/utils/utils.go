package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSONResponse writes a JSON response to the http.ResponseWriter.
func WriteJSONResponse(w http.ResponseWriter, jsonBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	// Marshal
	jsonBytes, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}
	// Write json bytes to string
	jsonString := string(jsonBytes)
	// Write string to response
	fmt.Fprint(w, jsonString)
}

func CheckValidRequest(r *http.Request) bool {
	ts := r.URL.Query().Get("ts")
	if ts == "" {
		fmt.Println("ts is empty")
		return false
	}
	sn := r.URL.Query().Get("sn")
	if sn == "" {
		fmt.Println("sn is empty")
		return false
	}
	sign := r.URL.Query().Get("sign")
	if sign == "" {
		fmt.Println("sign is empty")
		return false
	}
	// Calc md5(${sn}apiaccesskey||${ts}apiaccesskey||${route})
	route := r.URL.Path[len("/api/"):]
	hash := md5.Sum([]byte(sn + "apiaccesskey||" + ts + "apiaccesskey||" + route))
	validSign := hex.EncodeToString(hash[:])
	return validSign == sign
}
