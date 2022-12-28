package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	opts                                  = mqtt.NewClientOptions() //for mqtt client options
	client                                = mqtt.NewClient(opts)    //for mqtt client
	subTopic                              = "measurements/#"
	messagePubHandler mqtt.MessageHandler = func(subClient mqtt.Client, msg mqtt.Message) {
		log.Println(string(msg.Payload()))
	}
	connectHandler mqtt.OnConnectHandler = func(subClient mqtt.Client) {
		log.Println("Connected to MQTT")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(subClient mqtt.Client, err error) {
		log.Printf("Connect lost: %v", err)
	}
)

func connectMQTT() {
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
	client = mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	sub(subTopic)
}

func sub(subTopic string) {
	token := client.Subscribe(subTopic, 0, nil)
	token.Wait()
}

func Pub(pubTopic string, mes string) {
	token := client.Publish(pubTopic, 0, false, mes)
	token.Wait()
}

func MQTTSub() {
	connectMQTT()
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
		}
	}
}
