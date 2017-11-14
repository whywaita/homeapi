# Home API

[whywaita](https://github.com/whywaita)の家にある色んな物のラッパーAPI群です。

## インストール

`vars` パッケージに[IRKit](http://getirkit.com/)で用いる赤外線情報や認証情報があります。

```
package vars

const (
        JsonHomeLight = `{"format":"raw","freq":38,"data":[...]}` # 家の電気
        JsonAirConOn  = `{"format":"raw","freq":38,"data":[...]}` # エアコンの電源を付ける
        JsonAirConOff = `{"format":"raw","freq":38,"data":[...]}` # エアコンの電源を切る
)

const (
        ClientKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" # IRKitのclient key
        DeviceID  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" # IRKitのdevice id
)
```

IRKitのclient keyとdevice idは[こちら](http://getirkit.com/)の`IRKit Internet HTTP API`の項目を見て取得してください。

### ソースからのビルド

```
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ go build .
```

## 使い方

yayoiはサーバモードとCLIモードが存在しています。デフォルトではサーバモードで起動します。

### サーバモード

```
$ ./yayoi -m server
$ ./yayoi # どちらでも可能
```

### CLIモード

```
$ ./yayoi -m cli
```

電気を消す場合は以下のようにします。

```
$ ./yayoi -m cli
>>> irkit light off
HomeLight Off...
```
