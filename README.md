# weather-store

weather-storeは気象データを追加、取得するサービスです。  
気象データはPostgreSQLに蓄積されます。
また、weather-storeはgRPC通信のみ対応しています。

# Build Image

Docker でのビルドを想定しています。

```
$ git clone https://github.com/aki-yogiri/weather-store.git
$ cd weather-store
$ sudo docker build -t weather-store:v1.0.3 .
```

# Deploy on Kubernetes

```
$ kubectl apply -f <path>/<to>/<weather-store>/kubernetes/weather-store.yaml
```


# Configuration

weather-crawlerは以下の環境変数を利用します。

| variable | default | |
|----------|---------|-|
| DB_HOST | postgresql | PostgreSQLサーバのホスト名 |
| DB_PORT | 5432 | PostgreSQLサーバのポート番号 |
| DB_USER | none | PostgreSQLサーバのユーザ名 |
| DB_PASSWORD | none | DB_USERでログインするためのパスワード |
| DB_NAME | none | database名 |

