# mock_server

## how to use

プロジェクトルートで下記コマンドを実行する。

```shell
go run ./
```

実行後[localhost](http://localhost:8080/)にアクセスすれば確認が出来る。

## ファイルを利用する場合

プロジェクトルートにファイルを配置して

```go
b, err := ioutil.ReadFile("response.json")
```

を書き換える。
