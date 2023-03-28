# 1 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
# 2 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino" 2
# 3 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino" 2
// #include "millisDelay.h"
# 5 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino" 2
# 19 "/home/jabrix/projects/clone/xjx-biller-backend/devices/xjx_biller_devices/xjx_biller_devices.ino"
int states = 0x1; //
int pendingTimeout = 0; //

WiFiClient espClient;
PubSubClient client;
AsyncDelay relayTimeOut;

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
            digitalWrite(0, !0x0);

        if (received > 0)
            digitalWrite(0, !0x1);
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
    client.setServer("192.168.226.28" /* Server address*/, 1883);
    client.setCallback(callback);
    pinMode(0, 0x01);
    digitalWrite(0, !0x0);
    client.publish("/xjx/" + "client01" + "/ip", WiFi.localIP().toString().c_str(), true);
    // client.publish("/xjx/" + mqtt_client_id + "/", WiFi.localIP().toString().c_str(), true);
}

void setup_wifi()
{
    delay(10);
    Serial.println();
    Serial.print("Connecting to ");
    Serial.print("jabrix home" /* SSID*/);
    WiFi.begin("jabrix home" /* SSID*/, "mhafizha1011" /* password*/);
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
        if (client.connect(client, "jabrix", "12345"))
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
    client.subscribe("/xjx/client01/in" /* client subscribe*/);
    delay(500);
}
