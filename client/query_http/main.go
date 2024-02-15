package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb2 "github.com/bigrainnn/grpc_db_ops/internal/pb"
	"io/ioutil"
	"net/http"
)

func main() {
	req := &pb2.GetRecordReq{
		BusinessId:  "image_route",
		QueryFields: []string{"user_id", "route"},
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("JSON serialization error:", err)
		return
	}

	req2, err := http.NewRequest("POST", "http://localhost:8080/db_common_ops/getrecord", bytes.NewBuffer(jsonData))
	req2.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req2)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
