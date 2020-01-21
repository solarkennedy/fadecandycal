package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadecandycal/colors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func random(min, max int) uint8 {
	xr := rand.Intn(max-min) + min
	return uint8(xr)
}

func shouldIBeOn() bool {
	if isKodiPlayingVideo() {
		return false
	} else {
		pst, _ := time.LoadLocation("America/Los_Angeles")
		now := time.Now().In(pst)
		hour := now.Hour()
		fmt.Println(hour)
		return (hour >= 18 && hour <= 21) || (hour > 6 && hour <= 7)
	}
}

func displayPattern(oc *opc.Client, leds_len int) {
	m := opc.NewMessage(0)
	for i := 0; i < leds_len; i++ {
		m.SetLength(uint16(leds_len * 3))
		m.SetPixelColor(i, random(2, 255), random(2, 255), random(2, 255))
	}
	err := oc.Send(m)
	if err != nil {
		log.Println("couldn't send color", err)
	}

}

func turnOff(oc *opc.Client, leds_len int) {
	m := opc.NewMessage(0)
	for i := 0; i < leds_len; i++ {
		m.SetLength(uint16(leds_len * 3))
		m.SetPixelColor(i, 0, 0, 0)
	}
	err := oc.Send(m)
	if err != nil {
		log.Println("couldn't send color", err)
	}

}

func isKodiPlayingVideo() bool {
	players := getKodiGetActivePlayers()
	if len(players) >= 1 {
		player0 := players[0].(map[string]interface{})
		return player0["type"] == "video"
	} else {
		return false
	}
}

func getKodiGetActivePlayers() []interface{} {
	// curl -X POST -H "content-type:application/json" 'http://10.0.2.10:8080/jsonrpc' -d '{"jsonrpc": "2.0", "method": "Player.GetActivePlayers", "id": 1}'
	type Payload struct {
		Jsonrpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		ID      int    `json:"id"`
	}
	data := Payload{
		Jsonrpc: "2.0",
		Method:  "Player.GetActivePlayers",
		ID:      1,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://10.0.2.10:8080/jsonrpc", body)
	if err != nil {
		fmt.Println("Error talking to kodi: ")
		fmt.Println(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error talking to kodi: ")
		fmt.Println(err)
		return nil
	}
	var result map[string]interface{}
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body_bytes, &result)
	players := result["result"].([]interface{})
	defer resp.Body.Close()
	return players
}

func getOCClient() *opc.Client {
	server := "10.0.2.113:7890"
	oc := opc.NewClient()
	err := oc.Connect("tcp", server)
	if err != nil {
		log.Fatal("Could not connect to Fadecandy server", err)
	}
	return oc
}

func main() {
	leds_len := 50
	oc := getOCClient()

	for {
		color_pallete := colors.GetDaysColors(time.Now())
		fmt.Println(color_pallete)
		if shouldIBeOn() == true {
			fmt.Println("Should be on")
			displayPattern(oc, leds_len)
		} else {
			fmt.Println("Should be off")
			turnOff(oc, leds_len)
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}
