# heys-chat-server

Golang MQTT chat server

## Using
* EMQX(MQTT) Broker cluster (always subscribe)
* GO

## Architecture
* this server running for publish to MQTT broker.
* after publish, push message to MongoDB.
* subscribe MQTT Broker front-end app.