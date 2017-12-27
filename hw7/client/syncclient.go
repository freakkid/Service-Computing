package client

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"os"
)

// 
// q = ""
// from=en&to=zh&appid=20171226000109327&salt=pkMKVbG3W56acsG8JPJ7
func httpGetSync() {
	resp, err := http.Get("")
	if err != nil {
		fmt.Fprinln(os.Stderr,err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprinln(os.Stderr,err)
	}
	fmt.Println(string(body))
}

func randInt(min, max int) {
	rand.Seed(time.Now())
	return min + rand.Intn(max - min)
}