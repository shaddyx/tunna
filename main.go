package main

import (
"log"
"github.com/songgao/packets/ethernet"
"github.com/songgao/water"
	"os/exec"
)
//var log = logging.MustGetLogger("example")

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = "O_O"

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Switching interface up")
	cmd := exec.Command("ifconfig", config.Name, "up")
	cmd.Run()

	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())
	}
}
