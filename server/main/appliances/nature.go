package appliances

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type Nature struct {
	token string
}

func (n Nature) SetToken() {
	// tokenファイル読み込み
	bytes, err := ioutil.ReadFile("token")
	if err != nil {
		log.Fatal(err)
	}

	n.token = string(bytes)
}

// SendSignal idで指定した赤外線シグナルの送信
func (n Nature) SendSignal(id string) {
	url := "https://api.nature.global/1/signals/" + id + "/send"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+n.token)

	client := new(http.Client)
	log.Printf("Send request")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	log.Printf("%s", dumpResp)
}
