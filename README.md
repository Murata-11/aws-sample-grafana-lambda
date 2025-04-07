# 操作方法

以下のコマンドを使用して、main ファイルをビルドする。

```Bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
$ zip function.zip bootstrap
```
