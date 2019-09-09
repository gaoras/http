package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//Data ..
type Data struct {
	Args    string `json:"args"`
	Headers string `json:"headers"`
	Origin  string `json:"origin"`
	URL     string `json:"url"`
}

func main() {
	DigestTest()
	//header := httpGetUsingHeader()
	//jsonTest(header)
}

func makeJSON() {
	data := Data{Args: "123", Headers: "Accept", Origin: "12.34.56.789, 12.34.56.789", URL: "https://httpbin.org/get"}
	jsonbytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(jsonbytes))

}

func jsonTest(data []byte) {
	// JSONデコード
	var dst Data
	//var dst interface{}
	if err := json.Unmarshal(data, &dst); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	fmt.Println(dst)
}

func httpGetUsingHeader() []byte {
	uri := "https://httpbin.org/get"
	q := url.Values{}
	q.Add("query", "1234")
	req, _ := http.NewRequest("GET", uri+"?"+q.Encode(), nil)
	req.Header.Set("Authorization", "API_Key")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	return body
}

func simpleGet() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
