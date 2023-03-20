package config

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func MqttConfig() *mqtt.ClientOptions {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	uri := fmt.Sprintf("tcp://%s:%s", os.Getenv("MQTT_SERVER"), os.Getenv("MQTT_PORT"))
	if err != nil {
		log.Fatalf("Failed to parse %q broker address: %s", uri, err)
	}

	server := flag.String("server", uri, "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topic := flag.String("topic", os.Getenv("MQTT_TOPIC"), "Topic to subscribe to")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientid := flag.String("clientid", os.Getenv("MQTT_CLIENT_ID"), "A clientid for the connection")
	username := flag.String("username", os.Getenv("MQTT_USER"), "A username to authenticate to the MQTT server")
	password := flag.String("password", os.Getenv("MQTT_PASSWORD"), "Password to match username")
	flag.Parse()

	MqttConfig := mqtt.NewClientOptions()
	MqttConfig.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	MqttConfig.AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)
	MqttConfig.SetUsername(*username)
	MqttConfig.SetPassword(*password)
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	MqttConfig.SetTLSConfig(tlsConfig)
	return MqttConfig
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s Message: %s\n", message.Topic(), message.Payload())
}
