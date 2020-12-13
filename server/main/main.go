package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"

	"./appliances"
)

var appliancesHandler = new(appliances.Nature)

func main() {
	appliancesHandler.SetToken()

	addr := flag.String("a", "127.0.0.1", "IP Address")

	engine := gin.Default()
	engine.GET("/action", func(c *gin.Context) {
		// 制御を実行
		action()
		// 応答を返す
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	})

	engine.Run(*addr + ":55555")
}

/** JSONデコード用に構造体定義 */
type SignalList struct {
	ID string `json:"sigID"`
}

func action() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("SignalList.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	var siglist []SignalList
	if err := json.Unmarshal(bytes, &siglist); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを基にシグナルを順次送信
	for _, p := range siglist {
		appliancesHandler.SendSignal(p.ID)
		log.Printf("Sent signal : %s\n", p.ID)
	}

	// エアコンシグナル送信
	appliancesHandler.AirconSignalSend()

}
