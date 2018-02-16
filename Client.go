package main

import (
	"github.com/songgao/packets/ethernet"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	clientConfigFile = "clientConfig.yaml"
)


type ClientConfig struct {
	Interface struct {
		Name string
	}
	Server struct {
		Host string
		Port string
	}
}

func LoadClientConfig () (*ClientConfig, error){
	conf := &ClientConfig{}
	bytes, err := ioutil.ReadFile(clientConfigFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, conf)
	if err != nil {
		return nil, err
	}
	log.Println("Config is: ", conf)
	return conf, nil
}


func InitClient () error {
	clientConfig, err := LoadClientConfig()
	if err != nil {
		return err
	}
	ifce := InitIface(clientConfig.Interface.Name)

	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination() )
		log.Printf("Src: %s\n", frame.Source() )
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())
	}

}
