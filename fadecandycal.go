package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadecandycal/colors"
)

func Now() time.Time {
	pst, _ := time.LoadLocation("America/Los_Angeles")
	now := time.Now().In(pst)
	return now
}

func getEnvOverride() string {
	return os.Getenv("FADECANDYCAL_DATE")
}

func random(min, max int) uint8 {
	xr := rand.Intn(max-min) + min
	return uint8(xr)
}

func shouldIBeOn() bool {
	if getEnvOverride() != "" {
		return true
	} else if isKodiPlayingVideo() {
		return false
	} else {
		now := Now()
		hour := now.Hour()
		//		rise, set := getSunriseSunset()
		//return (now.After(set) && hour <= 21) || (now.After(rise) && hour <= 7)
		return (hour >= 18 && hour <= 21) || (hour > 6 && hour <= 7)
	}
}

func displayPattern(oc *opc.Client, leds_len int, color_palette []colors.Color) {
	m := opc.NewMessage(0)
	led_grouping := 1
	if len(color_palette) == 0 {
		for i := 0; i < leds_len; i++ {
			m.SetLength(uint16(leds_len * 3))
			m.SetPixelColor(i, random(2, 255), random(2, 255), random(2, 255))
		}
	} else {
		for i := 0; i < leds_len; i += led_grouping {
			c := color_palette[rand.Intn(len(color_palette))]
			for j := i; j < (i + led_grouping); j++ {
				m.SetLength(uint16(leds_len * 3))
				m.SetPixelColor(j, c.R, c.G, c.B)
				colors.PrintColorBlock(c)
			}
		}
	}
	err := oc.Send(m)
	fmt.Println()
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
		fmt.Println("Error talking to kodi: ", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error talking to kodi: ", err)
		return nil
	}
	var result map[string]interface{}
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body_bytes, &result)
	if err != nil {
		fmt.Println("Non fatal error parsing json from Kodi:", err)
	}
	players := result["result"].([]interface{})
	defer resp.Body.Close()
	return players
}

func getOCClient() *opc.Client {
	server := "fadecandycal:7890"
	oc := opc.NewClient()
	err := oc.Connect("tcp", server)
	if err != nil {
		log.Fatal("Could not connect to Fadecandy server", err)
	}
	return oc
}

func parseOverride(input string) time.Time {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	today := Now()
	parsed := time.Date(today.Year(), colors.MonthToMonth(month), day, 0, 0, 0, 0, today.Location())
	fmt.Printf("Parsed env override '%s' as '%s'\n", input, parsed)
	return parsed
}

func getToday() time.Time {
	override := getEnvOverride()
	if override != "" {
		return parseOverride(override)
	} else {
		return Now()
	}
}

func main() {
	leds_len := 64
	oc := getOCClient()

	for {
		today := getToday()
		color_palette := colors.GetDaysColors(today)
		if shouldIBeOn() {
			displayPattern(oc, leds_len, color_palette)
		} else {
			turnOff(oc, leds_len)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
