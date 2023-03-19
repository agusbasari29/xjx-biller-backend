package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func MqttConfig() *mqtt.ClientOptions {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	server := fmt.Sprintf("tcp://%s:%s", os.Getenv("MQTT_SERVER"), os.Getenv("MQTT_PORT"))
	uri, err := url.Parse(server)
	if err != nil {
		log.Fatalf("Failed to parse %q broker address: %s", server, err)
	}
	MqttConfig := &mqtt.ClientOptions{
		ClientID: os.Getenv("MQTT_CLIENT_ID"),
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
	}
	MqttConfig.Servers = append(MqttConfig.Servers, uri)
	return MqttConfig
}
