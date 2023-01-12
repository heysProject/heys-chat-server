const mqtt = require('mqtt');

// Web 의 경우 websocket 을 통한 tcp 연결을 해야 한다. 순수 tcp 연결은 불가능하다.
const addr = "mqtt://127.0.0.1";
const options = {
  // Clean session
  clean: true,
  connectTimeout: 4000,
  protocolId: 'MQTT',
  protocolVersion: 4,
  // Auth
  // 브로커가 기존 연결과 동일한 클라이언트 ID로 새 연결을 수신하면 기존 연결이 끊어져 EOF오류가 발생함
  clientId: "test_node_mqtt_client",
  username: "emqx_subscriber",
  password: "any",
  port: 1883,
};
const client = mqtt.connect(addr, options);
const topic = 'test/1/1';

client.on("connect", function () {
  console.log("connected");
  client.subscribe(topic, { qos: 1 }, function (err) {
    if (err) {
      console.error("not connected");
    } else {
      console.log("subscribed");
    }
  });

  let count = 0;
  while(count < 5) {
    setTimeout(() => {
      client.publish(topic, "Hello mqtt");
    }, count * 2000)
    count++;
  }

  if (count === 5) {
    setTimeout(() => {
      client.end();
    }, count * 2000)
  }
});

client.on("message", function (topic, message) {
  console.log("Received", topic, message.toString());
});
