package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	dac "go-http-digest-auth-client"
)

const (
	TestURI = "https://httpbin.org/digest-auth/auth/user/password/MD5"
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

func MyDigest() {
	resp := httpGetWithHeader(TestURI, "", "")
	res := resp.Header.Get("Www-Authenticate")
	//res := "Digest realm=\"me@kennethreitz.com\", nonce=\"c001b746b0ee4377de446700fcd0671b\", qop=\"auth\", opaque=\"47bab5a7dfde67ea3be182c308e62171\", algorithm=MD5, stale=FALSE"
	r := DecodeResponse(res)
	r.User = "user"
	r.Password = "password"
	tmp, _ := url.Parse(TestURI)
	r.URI = tmp.Path
	r.Response = CalcDigestResponse(&r)

	//Digest認証のヘッダ作成
	AuthValue := fmt.Sprintf("Digest username=\"%s\",realm=\"%s\",nonce=\"%s\",uri=\"%s\",cnonce=\"%s\",nc=%s,algorithm=%s,response=\"%s\",qop=%s,opaque=\"%s\"", r.User, r.Realm, r.Nonce, r.URI, r.CNonce, r.NC, r.Algorithm, r.Response, r.Qop, r.Opaque)

	resp2 := httpGetWithHeader(TestURI, "Authorization", AuthValue)
	fmt.Println(resp2.StatusCode)
}

func httpGetWithHeader(uri string, key, value string) *http.Response {
	req, _ := http.NewRequest("GET", uri, nil)
	if key != "" {
		req.Header.Set(key, value)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func DecodeResponse(res string) DigestAuthResponse {
	var ans DigestAuthResponse

	resSplit := strings.Split(res, ",")

	ans.Realm = TrimString(resSplit, "realm")
	ans.Nonce = TrimString(resSplit, "nonce")
	ans.Qop = TrimString(resSplit, "qop")
	ans.Opaque = TrimString(resSplit, "opaque")
	ans.Algorithm = TrimString(resSplit, "algorithm")
	return ans
}

func TrimString(st []string, key string) string {
	var ans string
	for i := 0; i < len(st); i++ {
		n := strings.Index(st[i], key)
		if n >= 0 {
			trim := strings.Trim(st[i], "Digest ")
			trim2 := strings.Trim(trim, key)
			trim3 := strings.Trim(trim2, " \"=")
			ans = trim3
			break
		}
	}
	return ans
}

func CalcDigestResponse(res *DigestAuthResponse) string {
	user := res.User
	password := res.Password
	realm := res.Realm
	method := "GET"
	URI := res.URI

	res.NC = "00000001" //本来は内部でカウントアップしていく
	res.CNonce = "3f5289181d628df1bab9111ad2c7e40f"
	A1 := fmt.Sprintf("%s:%s:%s", user, realm, password)
	A2 := fmt.Sprintf("%s:%s", method, URI)

	respons := MD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", MD5(A1), res.Nonce, res.NC, res.CNonce, res.Qop, MD5(A2)))
	return respons
}

func MD5(st string) string {
	ans := fmt.Sprintf("%x", md5.Sum([]byte(st)))
	return ans
}

type DigestAuthResponse struct {
	User      string
	Password  string
	URI       string
	Realm     string
	Nonce     string
	CNonce    string
	NC        string
	Qop       string
	Opaque    string
	Algorithm string
	Stale     bool
	Response  string
}
