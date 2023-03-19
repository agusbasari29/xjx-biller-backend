package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/agusbasari29/xjx-biller-backend/database"
	"github.com/agusbasari29/xjx-biller-backend/entity"
	"github.com/agusbasari29/xjx-biller-backend/queue"
	"github.com/agusbasari29/xjx-biller-backend/tasks"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB    = database.SetupDatabaseConnection()
	client mqtt.Client = queue.SetupMqttConnection()
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	defer database.CloseDatabaseConnection(db)
	// subscribe := tasks.NewSubscribeTasks(client)
	pub := tasks.NewPublishTask(client)
	// callbacks := tasks.onMessageReceived(client)
	db.AutoMigrate(&entity.Users{}, &entity.Clients{}, &entity.Products{}, &entity.Transaction{}, &entity.ItemsTrx{})
	g := gin.Default()
	g.Run(os.Getenv("SERVER_PORT"))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// subscribe.Subscribe("/xjx/#")
	pub.Publish("test", "/xjx/test")
	<-c
}
