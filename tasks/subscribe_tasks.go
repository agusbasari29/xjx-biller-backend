package tasks

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeTasks interface {
	Subscribe(topic string)
	// onMessageReceived(client mqtt.Client, msg mqtt.Message)
}

type subscribeTasks struct {
	client mqtt.Client
}

func NewSubscribeTasks(client mqtt.Client) *subscribeTasks {
	return &subscribeTasks{client}
}

func (s *subscribeTasks) Subscribe(topic string) {
	token := s.client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

// func (s *subscribeTasks) onMessageReceived(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message %s from topic: %s\n", msg.Payload(), msg.Topic())
// }
