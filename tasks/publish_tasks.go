package tasks

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type PublishTask interface {
	Publish(msg string, topic string)
}

type publishTask struct {
	client mqtt.Client
}

func NewPublishTask(client mqtt.Client) *publishTask {
	return &publishTask{client}
}

func (p *publishTask) Publish(msg string, topic string) {
	text := fmt.Sprintf("Message %s", msg)
	token := p.client.Publish(topic, 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}
