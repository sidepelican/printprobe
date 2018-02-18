# PrintProbe
プローブ要求をキャプチャしてみようという課題

Raspberry PiやMacでプローブ要求をキャプチャして内容を表示するアプリケーションを作ってみよう💪


## Go環境のセットアップ
サンプルプログラムはgolangで書かれています
- Goのインストール手順: http://golang-jp.org/doc/install

golangで作成したプログラムは1つのバイナルファイルとなるため取り回しが容易です
- Node.jsやRubyでは `npm install` などが必要で、特にセンサ環境ではインストール時にコケることが多いです
- また `node` のバージョンが違うなどのエラーが出るのは日常茶飯事であり、それらの厳密な管理も必要となります
- 貧弱なメモリ・プロセッサの場合はインストールに時間がかかることもしばしばです

このため複数センサ環境においてインストールが容易でどのセンサ端末に置いても同じ動作が期待できるシングルバイナリ型の実行ファイルが有用であると考え、
そのためにGoが用いられています


## プログラムのコンパイル＆実行

このサンプルプログラムでは外部パッケージとして `github.com/sidepelican/goprobe/probe` ([goprobe](https://github.com/sidepelican/goprobe))を使用しています  
コンパイルする前にあらかじめ使うパッケージを `go get` する必要があります

```
> go get github.com/sidepelican/goprobe/probe
```

ビルドは `go build` コマンドです。自動的にコンパイル対象ファイルを探して名前もつけてくれます

```
> go build
```

ネットワークインターフェースをモニタモードで使用するためには管理者権限が必要となるため、実行する際はsudoをつけます

```
> sudo ./printprobe
2018/02/18 22:01:20 used interface: en1
2018/02/18 22:01:20 pcap version: libpcap version 1.8.1 -- Apple version 79.20.1
...
```

## 取得したデータを転送する

どうやって転送するか🤔  
👉 MQTT  
👉 HTTP POST  
👉 fluentd

### MQTTの例
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

client := MQTT.NewClient(opts)
if token := client.Connect(); token.Wait() && token.Error() != nil {
    fmt.Println("MQTT Error:", token.Error())
    return
}
defer client.Disconnect(250)

...

for record := range source.Records() {
    ...

    // MQTTで送信
    payload := []byte(record.String())
    client.Publish("your/topic", 2, false, payload)
}
```
