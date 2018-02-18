# PrintProbe
プローブ要求をキャプチャしてみようという課題

Raspberry PiやMacでプローブ要求をキャプチャして内容を表示するアプリケーションを作ってみよう💪


## Go環境のセットアップ
サンプルプログラムはgolangで書かれています。
- goのインストール手順: http://golang-jp.org/doc/install

## プログラムのコンパイル＆実行

```
> go build
> sudo ./printprobe
2018/02/18 22:01:20 used interface: en1
2018/02/18 22:01:20 pcap version: libpcap version 1.8.1 -- Apple version 79.20.1
...
```

ネットワークインターフェースをモニタモードで使用するためには管理者権限が必要となるため、sudoなどをつけないとエラーになります

## 取得したデータを転送する

どうやって転送するか🤔  
👉 MQTT  
👉 HTTP POST  
👉 fluentd

# MQTTの例
MQTTのライブラリ `phao.mqtt` を使用する

```:go
import (
    ...

    MQTT "github.com/eclipse/paho.mqtt.golang" // 👈 追加
)
```

```:go
// MQTTクライアントのセットアップ
opts := MQTT.NewClientOptions()
opts.AddBroker("your.broker.address")

mqttClient := MQTT.NewClient(opts)
if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
    fmt.Println("MQTT Error:", token.Error())
    return
}
defer mqttClient.Disconnect(250)

...

for record := range source.Records() {
    ...

    // MQTTで送信
    payload := []byte(record.String())
    client.Publish("your/topic", 2, false, payload)
}
```
