package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")

	// ファイルのレスポンスを返す場合
	//b, err := ioutil.ReadFile("response.json")
	//if err != nil {
	//	panic(err)
	//}
	//json := string(b)
	//fmt.Fprintf(w, json)

}
