#include <Arduino.h>
#line 1 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
#include <ESP8266WiFi.h>
#include <PubSubClient.h>
// #include "millisDelay.h"
#include <AsyncDelay.h>

#define wifi_ssid "jabrix home"      // SSID
#define wifi_password "mhafizha1011" // password

#define mqtt_server "192.168.226.28" // Server address
#define mqtt_port 1883
#define mqtt_user "jabrix"
#define mqtt_password "12345"
#define mqtt_client_id "client01"
#define in_topic "/xjx/client01/in"   // client subscribe
#define out_topic "/xjx/client01/out" // client publish

#define relay 0

int states = HIGH;      //
int pendingTimeout = 0; //

WiFiClient espClient;
PubSubClient client;
AsyncDelay relayTimeOut;

#line 26 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void callback(char *topic, byte *payload, unsigned int length);
#line 47 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void setup();
#line 61 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void setup_wifi();
#line 79 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void reconnect();
#line 98 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void loop();
#line 26 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
void callback(char *topic, byte *payload, unsigned int length)
{
    Serial.print("Message arrived [");
    Serial.print(topic);
    Serial.print("] ");
    for (int i = 0; i < length; i++)
    {
        char receivedChar = (char)payload[i];
        Serial.print(receivedChar);
        char *c = receivedChar;
        received = atoi(receivedChar);
        if (received == 0)
            digitalWrite(relay, !LOW);

        if (received > 0)
            digitalWrite(relay, !HIGH);
        timeOut.start();
    }
    Serial.println(atoi().toInt());
}

void setup()
{
    // put your setup code here, to run once:
    Serial.begin(115200);
    setup_wifi();
    client.setClient(espClient);
    client.setServer(mqtt_server, mqtt_port);
    client.setCallback(callback);
    pinMode(relay, OUTPUT);
    digitalWrite(relay, !LOW);
    client.publish("/xjx/" + mqtt_client_id + "/ip", WiFi.localIP().toString().c_str(), true);
    // client.publish("/xjx/" + mqtt_client_id + "/", WiFi.localIP().toString().c_str(), true);
}

void setup_wifi()
{
    delay(10);
    Serial.println();
    Serial.print("Connecting to ");
    Serial.print(wifi_ssid);
    WiFi.begin(wifi_ssid, wifi_password);
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(500);
        Serial.print(".");
    }
    Serial.println("");
    Serial.println("WiFi connected");
    Serial.println("IP Address: ");
    Serial.println(WiFi.localIP());
}

void reconnect()
{
    while (!client.connected())
    {
        Serial.print("Attemping MQTT connection...");
        if (client.connect(client, mqtt_user, mqtt_password))
        {
            Serial.println("connected");
        }
        else
        {
            Serial.print("failed, rc=");
            Serial.print(client.state());
            Serial.println(" try again in 5 seconds");
            delay(5000);
        }
    }
}

void loop()
{
    if (!client.connected())
    {
        reconnect();
    }
    client.loop();
    client.subscribe(in_topic);
    delay(500);
}

