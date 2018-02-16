package main

import (
	"github.com/songgao/water"
	"os/exec"
	"log"
)

func InitIface(name string) *water.Interface{
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = name

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Switching interface up")
	cmd := exec.Command("ifconfig", config.Name, "up")
	cmd.Run()
	return ifce
}
