package appliances

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// SendSignal idで指定した赤外線シグナルの送信
func SendSignal(id string) {
	url := "https://api.nature.global/1/signals/" + id + "/send"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+token)

	client := new(http.Client)
	log.Printf("Send request")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	log.Printf("%s", dumpResp)
}
