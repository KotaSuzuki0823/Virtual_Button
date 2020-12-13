package appliances

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type Nature struct {
	// Nature Remoのトークン
	token string

	// エアコン用変数
	// ID
	airconID string
	// 温度
	airconTemperature string
	// 運転モード
	airconMode string
	// 風量
	airconAirVolume string
	// 風向
	airconAirDirection string
}

/** JSONデコード用に構造体定義 */
type temp struct {
	// エアコン用変数
	// ID
	id string `json:"sigID"`
	// 温度
	temperature string `json:"temperature"`
	// 運転モード
	mode string `json:"mode"`
	// 風量
	airVolume string `json:"volume"`
	// 風向
	airDirection string `json:"direction"`
}

// SetToken tokenファイルから読み込み
func (n Nature) SetToken() {
	// tokenファイル読み込み
	n.token = os.Getenv("NATURE_TOKEN")
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

// loadAirconSettings AirconSettings.jsonから送信したいエアコンシグナルを読み込み
func (n Nature) loadAirconSettings() bool {
	// jsonファイル読み込み
	bytes, err := ioutil.ReadFile("AirconSettings.json")
	if err != nil {
		log.Print(err)
		log.Print("AirconSettings.jsonが見つからないため，エアコン操作は行われません．")
		return false
	}

	// JSONデコード
	var list []temp
	if err := json.Unmarshal(bytes, &list); err != nil {
		log.Print(err)
		log.Print("AirconSettings.jsonの解析に失敗したため，エアコン操作は行われません．")
		return false
	}

	n.airconID = list[0].id
	n.airconTemperature = list[0].temperature
	n.airconMode = list[0].mode
	n.airconAirVolume = list[0].airVolume
	n.airconAirDirection = list[0].airDirection

	return true
}

// AirconSignalSend エアコン用シグナル送信
func (n Nature) AirconSignalSend() {
	if n.loadAirconSettings() {
		remourl := "https://api.nature.global/1/appliances/" + n.airconID + "/aircon_settings"
		value := url.Values{}
		// クエリ追加
		value.Add("temperature", n.airconTemperature)
		value.Add("operation_mode", n.airconMode)
		value.Add("air_volume", n.airconAirVolume)
		value.Add("air_direction", n.airconAirDirection)
		// 電源ONのため，空欄
		value.Add("button", "")
		// 電源OFFの場合
		//value.Add("button", "poewr-off")

		req, _ := http.NewRequest(
			"POST",
			remourl,
			strings.NewReader(value.Encode()),
		)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", " Bearer "+n.token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
	} else {
		// AirconSettings.jsonがない場合，何もしない
		return
	}
}
