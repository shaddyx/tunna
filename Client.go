package main

import (
	"github.com/songgao/packets/ethernet"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)


type ClientConfig struct {
	Interface struct {
		Name string
	}
	Server struct {
		Host string
		Port int
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

func initWsClient(conf ClientConfig) (*gosocketio.Client, error){
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(conf.Server.Host, conf.Server.Port, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil{
		return nil, err
	}
	return c, nil
}


func InitClient () error {
	clientConfig, err := LoadClientConfig()
	if err != nil {
		return err
	}
	ifce := InitIface(clientConfig.Interface.Name)
	log.Println("Connecting to websocket server...")
	wsClient, err:= initWsClient(*clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Joining room...")
	wsClient.Join(Room)
	log.Println("adding listener...")
	wsClient.On(DataEvent, func(c *gosocketio.Channel, msg DataGram) string {
		return "OK"
	})

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
