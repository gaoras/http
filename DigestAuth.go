package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	dac "go-http-digest-auth-client"
)

const (
	TestURI = "https://httpbin.org/digest-auth/qop/user/password/MD5"
)

func DigestTest() {
	fmt.Println("Digest Test")

	t := dac.NewTransport("user", "password")
	req, err := http.NewRequest("GET", TestURI, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
