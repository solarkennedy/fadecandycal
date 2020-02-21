package main

import (
	_ "fmt"
	"log"

	"github.com/kellydunn/go-opc"
)

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
	ledStart := 40
	ledEnd := 50
	oc := getOCClient()
	m := opc.NewMessage(0)
	for i := ledStart; i < ledEnd; i++ {
		m.SetLength(uint16(ledStart - ledEnd*3))
		m.SetPixelColor(i, 255, 0, 0)
	}
	err := oc.Send(m)
	if err != nil {
		log.Println("couldn't send color", err)
	}

}
