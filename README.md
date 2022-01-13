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