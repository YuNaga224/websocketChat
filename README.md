## websocketを使ったチャットアプリです
Go言語で実装しています。「Go言語によるWebアプリケーション開発」という書籍を参考に作りました。

[Go言語によるwebアプリケーション開発（amazon ）](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526/ref=sr_1_1?adgrpid=53112876877&gclid=CjwKCAiA_vKeBhAdEiwAFb_nraMDYWr9_2_zAsrr3dOx5zlwshxpVAYRBzeyNlsdsY7Y1jfmYsOWlBoCGWQQAvD_BwE&hvadid=618553090956&hvdev=c&hvlocphy=1009076&hvnetw=g&hvqmt=e&hvrand=13216602359019540869&hvtargid=kwd-335468543278&hydadcr=27294_14598076&jp-ad-ap=0&keywords=go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8Bweb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA&qid=1675413707&sr=8-1)

## 使い方

まずコマンドラインで以下を実行してください
```
git clone https://github.com/YuNaga224/GoWebsocketChat.git

go mod tidy

```
次にルートディレクトリにconfig.jsonを以下のような形式で作成してください。
```
{
    "google":{
        "client_id":"クライアントID",
        "client_secret":"クライアントシークレット"
    },
    "github":{
        "client_id":"クライアントID",
        "client_secret":"クライアントシークレット"
    }
}
```
その後コマンドラインで以下を実行するとサーバーが立ち上がります
```
go build -o chat

./chat
```

localhost:8080にアクセスすると使えます。
実際に複数のブラウザを立ち上げてチャットしてみてください。
ctrl + Cで終了できます。


