# 操作方法

以下のコマンドを使用して、main ファイルをビルドする。

```Bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
$ zip function.zip bootstrap
```

## Grafana 参考 URL

- [Grafana ダッシュボード作成ハンズオン](https://github.com/classmethod/grafana_dashboard_handson/blob/main/docs/handson-guide.md)
- [Postman モックサーバーを使って Grafana（Infinity プラグイン）の動作検証を行う](https://qiita.com/yankee/items/69e5fb4c5b1d79606b35)
- [(YouTube) Grafana Dashboard using JSON Data and Grafana Inifinity Plugin](https://www.youtube.com/watch?v=vAoR83g6CE4&t=299s)
