package queue

import (
	"fmt"

	"github.com/agusbasari29/xjx-biller-backend/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
// }

// var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
// 	fmt.Printf("Connected to")
// 	// if token := client.Subscribe("/xjx/#", 0, messagePubHandler); token.Wait() && token.Error() != nil {
// 	// 	panic(token.Error())
// 	// }
// }

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func SetupMqttConnection() mqtt.Client {
	mqttConfig := config.MqttConfig()
	// mqttConfig.SetDefaultPublishHandler(messagePubHandler)
	mqttConfig.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe("/xjx/#", 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	mqttConfig.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(mqttConfig)
	return client
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s Message: %s\n", message.Topic(), message.Payload())
}
