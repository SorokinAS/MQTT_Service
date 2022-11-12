package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	opts                                  = mqtt.NewClientOptions() //for mqtt subClient options
	subClient                             = mqtt.NewClient(opts)    //for mqtt subClient
	topic                                 = "measurements/#"
	messagePubHandler mqtt.MessageHandler = func(subClient mqtt.Client, msg mqtt.Message) {
		fmt.Println(string(msg.Payload()))
	}
	connectHandler mqtt.OnConnectHandler = func(subClient mqtt.Client) {
		fmt.Println("Connected")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(subClient mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
)

func ConnectMQTT() {
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetKeepAlive(15 * time.Second)
	opts.SetCleanSession(true)
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(15 * time.Second)
	opts.SetAutoReconnect(true).SetMaxReconnectInterval(15 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	subClient = mqtt.NewClient(opts)
	token := subClient.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(topic)
}

func sub(topic string) {
	token := subClient.Subscribe(topic, 0, nil)
	token.Wait()
}

func main() {
	ConnectMQTT()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
		}
	}
}
