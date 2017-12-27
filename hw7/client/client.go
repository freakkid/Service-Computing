package client

import (
	//"strings"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
	"strconv"
)

const (
	min = 100000
	max = 999999
	from = "en"
	to = "zh"
	appid = "20171226000109327"
	key = "pkMKVbG3W56acsG8JPJ7"
)

type resultJSON struct {
	From string `json:"from"`
	To string `json:"to"`
	TransResult []map[string]string `json:"trans_result"`
}

// HTTPGet .
func HTTPGet(q string) {
	resp, err := http.Get(generateURL(q))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var resultjson resultJSON
	if err = json.Unmarshal(body, &resultjson); err == nil {
		fmt.Println(resultjson.TransResult[0]["src"], resultjson.TransResult[0]["dst"])
	} else {
		fmt.Println(q + ": Ocurr error when parsing json response")
	}
}

// HTTPGetWithChannel .
func HTTPGetWithChannel(q string, s chan string) {
	resp, err := http.Get(generateURL(q))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var resultjson resultJSON
	if err = json.Unmarshal(body, &resultjson); err == nil {
		s <-fmt.Sprintln(resultjson.TransResult[0]["src"], resultjson.TransResult[0]["dst"]) 
	} else {
		fmt.Println(q + ": Ocurr error when parsing json response")
		s <- ""
	}
}

// from=en&to=zh&appid=20171226000109327&key=pkMKVbG3W56acsG8JPJ7
func generateURL(q string) string {
	salt := generateSalt()
	sign := generateSign(q, salt)
	return "http://api.fanyi.baidu.com/api/trans/vip/translate?q="+q+"&from="+from+"&to="+to+"&appid=20171226000109327&salt="+salt+"&sign="+sign
}

// appid + q + salt + key
func generateSign(q string, salt string) string {
	return encryptByMD5("20171226000109327"+ q + salt + key)
}

func generateSalt() string {
	return getRandInt(min, max)
}

func getRandInt(min, max int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min))
}
func encryptByMD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}