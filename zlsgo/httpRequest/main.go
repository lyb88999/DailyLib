package main

import (
	"fmt"
	"github.com/sohaha/zlsgo/zhttp"
)

func main() {
	// Get
	urlGet := "https://httpbin.org/get"
	GetHB(urlGet)

	urlPost := "https://httpbin.org/post"

	PostHB(urlPost, zhttp.BodyJSON(map[string]string{
		"name": "lyb",
		"age":  "18",
	}))

	urlPut := "https://httpbin.org/put"
	PutHB(urlPut, zhttp.BodyJSON(map[string]string{
		"name": "lyb123",
		"age":  "23",
	}))

	//urlUpload := "https://httpbin.org/post"
	//UploadHB(urlUpload, "/Users/lyb/GolandProjects/DailyLib/zlsgo/httpRequest/hiking.png")
	//
	//urlDownload := "https://httpbin.org/bytes/104857600"
	//DownloadHB(urlDownload)
}

func GetHB(url string) {
	data, err := zhttp.Get(url)
	if err != nil {
		panic(err)
	}
	res := data.JSONs()
	fmt.Println(res)
}

func PostHB(url string, params ...interface{}) {
	data, err := zhttp.Post(url, params...)
	if err != nil {
		panic(err)
	}
	res := data.JSONs()
	fmt.Println(res)
}

func PutHB(url string, params ...interface{}) {
	data, err := zhttp.Put(url, params...)
	if err != nil {
		panic(err)
	}
	res := data.JSONs()
	fmt.Println(res)
}

func UploadHB(url string, filePath string) {
	data, err := zhttp.Post(url, zhttp.File(filePath))
	if err != nil {
		panic(err)
	}
	res := data.JSONs()
	fmt.Println(res)
}

func DownloadHB(url string) {
	data, err := zhttp.Get(url)
	if err != nil {
		panic(err)
	}
	res := data.Bytes()
	fmt.Println(res)
}
