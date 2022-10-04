package tavanir

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"tools"
)

// APIRequest send API request
func APIRequest(methodName string, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", baseURL+methodName, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Http request builder failed with error:\n%s\n", err)
	}
	//log.Println("methodName: ", methodName )
	//log.Println("token: ", ComputeAuthToken(methodName) )
	req.Header.Set("token", ComputeAuthToken(methodName))
	req.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		fmt.Printf("[HTTP error] %d : %s\n", response.StatusCode, baseURL+methodName)
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		return nil, fmt.Errorf("HTTP error: [%d] %s", response.StatusCode, string(data))
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	return data, nil
}

func makeTimestamp() int64 {
	return time.Now().Unix()
}

func computeHmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func computeAuthToken(apiName, methodName string) string {
	secret := tools.GetEnv("API_SECRET", "123456")
	timestamp := makeTimestamp()
	return fmt.Sprintf("%d-%s", timestamp, computeHmac256(fmt.Sprintf("%s.%s_%d", apiName, methodName, timestamp), secret))
}

//ComputeAuthToken generate auth token
func ComputeAuthToken(methodName string) string {
	return computeAuthToken("Insurance", methodName)
}
