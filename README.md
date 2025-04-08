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

## QuickSight Athena DynamoDB

- [QuickSight で DynamoDB のデータを可視化するためのフェデレーテッドクエリ用 Lambda を選択できない](https://community.amazonquicksight.com/t/quicksight-dynamodb-lambda/44402)

QuickSight のデフォルトロールに以下のポリシーを追加する

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "lambda:InvokeFunction",
      "Resource": "arn:aws:lambda:<リージョン>:<アカウントID>:function:<Lambda関数名>"
    }
  ]
}
```

QuickSight からの SQL 例

```sql
SELECT
    pk,
    format_datetime(from_iso8601_timestamp(sk), 'yyyy-MM-dd HH:mm:ss') as sk,
    value
FROM (
    SELECT *, ROW_NUMBER() OVER (PARTITION BY pk ORDER BY sk DESC) AS rn
    FROM "Sample-QuickSight-Athena"."default"."sample-quicksight-athena"
) AS ranked_data
WHERE rn = 1
```
