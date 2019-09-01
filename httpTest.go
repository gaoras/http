package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	httpGetUsingHeader()
}

func httpGetUsingHeader() {
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
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	str := string(robots)
	spStr := strings.Split(str, "\n")
	fmt.Println(spStr)
	for index := 0; index < len(spStr); index++ {
		n := strings.Index(spStr[index], "Authorization")
		if n > 0 {
			target := spStr[index][n:]
			fmt.Println(target)
			break
		}
	}
	//	fmt.Println(n)
}

func simpleGet() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
