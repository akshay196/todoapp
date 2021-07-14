package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var SERVERIP = "http://localhost:8000"

func main() {
	resp, err := http.Get(SERVERIP + "/list")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("/list returned status not OK")
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return
	}
	list := make(map[int]string)
	err = json.Unmarshal(b, &list)
	if err != nil {
		fmt.Println("Error parsing the JSON data")
		return
	}
	fmt.Println(list)
}
