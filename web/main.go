package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello Bitches!!!")

	// PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	//! here we get the response body in byte format
	content, _ := ioutil.ReadAll(res.Body)

	var responseString strings.Builder
	responseString.Write(content)

	fmt.Println("Content:", string(content))
	fmt.Println("Content:", responseString.String())
}

func PerformPostRequest() {
	const url = "http://localhost:8000/post"

	//fake json data
	jsonData := `{"name":"John", "age":30}`

	res, err := http.Post(url, "application/json", strings.NewReader(jsonData))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	var responseString strings.Builder
	responseString.Write(content)

	fmt.Println("Content:", responseString.String())
}
