package mqtt

import (
	"log"
	"os"
	"regexp"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Broker struct {
	Host     string
	SubTopic string
}

var (
	opts              = mqtt.NewClientOptions() //for mqtt client options
	client            = mqtt.NewClient(opts)    //for mqtt client
	meas              string
	topicName         string
	messagePubHandler mqtt.MessageHandler = func(subClient mqtt.Client, msg mqtt.Message) {
		topicName = msg.Topic()
		meas = string(msg.Payload())
		changeString(&meas, *&reg)
		sendMeasurements(&meas, &topicName)
	}
	connectHandler mqtt.OnConnectHandler = func(subClient mqtt.Client) {
		log.Println("Connected to MQTT")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(subClient mqtt.Client, err error) {
		log.Fatalf("Connect lost: %v", err)
	}
	reg = regexp.MustCompile(`\[|\]`)
)

func connectMQTT() {
	broker := Broker{
		Host:     os.Getenv("MQTT_BROKER_URL"),
		SubTopic: os.Getenv("MQTT_SUB_TOPIC"),
	}
	opts.AddBroker(broker.Host)
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
	sub(broker.SubTopic)
}

func sub(subTopic string) {
	token := client.Subscribe(subTopic, 0, nil)
	token.Wait()
}

func Pub(pubTopic string, mes string) {
	changeString(&mes, *&reg)
	log.Print(mes)
	token := client.Publish(pubTopic, 0, false, mes)
	token.Wait()
}

func Run() {
	connectMQTT()
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
		}
	}
}

func changeString(s *string, re *regexp.Regexp) {
	*s = string(re.ReplaceAll([]byte(*s), []byte("")))
}
