# Go-API-test

## using

- 動作確認
~~~bash
curl localhost:3000/
// show test massage
~~~

- データ表示（追加後は増加する）
~~~bash
curl localhost:3000/show
// 以下のような初期データを表示
    {
        "name": "maruto",
        "pass": "maruto"
    }
~~~
- データ追加
~~~bash
curl -X POST -H "content-type:application/json" localhost:3000/post -d '{"name": "a","pass":"aa"}'
// データの追加
~~~

[GinとAPIに関する参考資料](https://go.dev/doc/tutorial/web-service-gin)

[curlに関する参考資料](https://qiita.com/yasuhiroki/items/a569d3371a66e365316f)

[環境構築関する参考資料](https://www.utakata.work/entry//golang/tutorial/2-docker-go-gin)