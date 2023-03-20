package queue

import (
	"github.com/agusbasari29/xjx-biller-backend/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SetupMqttConnection() mqtt.Client {
	mqttConfig := config.MqttConfig()
	client := mqtt.NewClient(mqttConfig)
	return client
}
