package adapter

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"rubix-lib-rest-go/config"
	"rubix-lib-rest-go/controller"
)

type MqttConnection struct {
	mqttClient mqtt.Client
}

func NewConnection(clientId string) (conn *MqttConnection) {
	c := config.CommonConfig()
	opts := mqtt.NewClientOptions()
	host := "tcp://" + c.Broker.Host + ":" + c.Broker.Port
	opts.AddBroker(fmt.Sprintf(host))
	opts.SetClientID(clientId)
	opts.AutoReconnect = true
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("Connect problem: ", token.Error())
	}
	conn = &MqttConnection{client}
	return conn
}

func (conn *MqttConnection) Subscribe(topic string) {
	token := conn.mqttClient.Subscribe(topic, 1, onMessageReceived())
	token.Wait()
	log.Println("Subscribed to topic: ", topic)
}

func (con *MqttConnection) IsConnected() bool {
	connected := con.mqttClient.IsConnected()
	if !connected {
		log.Println("Healthcheck MQTT fails")
	}
	return connected
}

func (conn *MqttConnection) Publish(message string, topic string) {
	token := conn.mqttClient.Publish(topic, 1, false, message)
	token.Wait()
	log.Println("Publish to topic: ", topic)
}

func onMessageReceived() func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
		event := controller.ChipEvent{}
		err := json.Unmarshal([]byte(msg.Payload()), &event)
		if err != nil {
			log.Println("Unmarshal message fails: ", err)
		}

	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt connected")
}
