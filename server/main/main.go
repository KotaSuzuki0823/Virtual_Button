package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"./appliances"
)

var appliancesHandler = new(appliances.Nature)

func main() {
	appliancesHandler.SetToken()

	addr := flag.String("a", "127.0.0.1", "IP Address")
	flag.Parse()

	engine := gin.Default()
	engine.GET("/action", func(c *gin.Context) {
		// 制御を実行
		if err := action(); err != nil {
			log.Print(err)
			// 応答を返す
			c.JSON(http.StatusOK, gin.H{
				"message": "fault",
			})
		} else {
			// 応答を返す
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
	})

	engine.Run(*addr + ":55555")
}

/** JSONデコード用に構造体定義 */
type SignalList struct {
	ID string `json:"sigID"`
}

func action() error {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./settings/SignalList.json")
	if err != nil {
		return err
	}

	// JSONデコード
	var siglist []SignalList
	if err := json.Unmarshal(bytes, &siglist); err != nil {
		return err
	}
	// デコードしたデータを基にシグナルを順次送信
	for _, p := range siglist {
		appliancesHandler.SendSignal(p.ID)
		log.Printf("Sent signal : %s\n", p.ID)
	}

	// エアコンシグナル送信
	appliancesHandler.AirconSignalSend()

	return nil
}
