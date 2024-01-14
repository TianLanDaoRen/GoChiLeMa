package route_utils

import (
	"GoChiLeMaWails/src/encrypto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonBody interface{}

type RequestBody struct {
	Sign string  `json:"sign"`
	Sn   float64 `json:"sn"`
	Ts   float64 `json:"ts"`
}

func MakeDefaultJSONResponse(err int, msg string) map[string]interface{} {
	jsonResponse := make(map[string]interface{})
	jsonResponse["error"] = err
	jsonResponse["message"] = msg
	return jsonResponse
}

// WriteJSONResponse writes a JSON response to the http.ResponseWriter.
func WriteJSONResponse(w http.ResponseWriter, jsonBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
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

func ReadJsonBody(r *http.Request, jb JsonBody) error {
	// Check request method
	if r.Method != "POST" {
		return fmt.Errorf("INVALID REQUEST METHOD: %v", r.Method)
	}
	// Check request content type
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("INVALID REQUEST CONTENT-TYPE: %v", r.Header.Get("Content-Type"))
	}
	bodyBytes, err := io.ReadAll(r.Body)
	// Restore request body
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	// Get request body
	err = json.Unmarshal(bodyBytes, jb)
	if err != nil {
		return err
	}
	return nil
}

func CheckValidRequest(r *http.Request) bool {
	// Get request body
	var body RequestBody
	err := ReadJsonBody(r, &body)
	if err != nil {
		fmt.Println("Failed to read request body:", err)
		return false
	}
	// Check valid sign
	// Encrypt(${sn}apiaccesskey||${ts}apiaccesskey||${route})
	crypto := encrypto.NewEncrypto()
	route := r.URL.Path[len("/api/"):]
	ciphertext := fmt.Sprintf("%vapiaccesskey||%vapiaccesskey||%v", body.Sn, int(body.Ts), route)
	validSign := crypto.Encrypt(ciphertext)
	return validSign == body.Sign
}
