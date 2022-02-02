# Go-API-tes
## deployment
- dockerコンテナ上で動作確認

    `docker-compose up`
- kubernatesでpodを立てて動作確認

    ```zh
    docker-compose build -t <image名(test-k8s)> .
    kubectl apply -f deployment.yml
    kubectl apply -f service.yml
    ```
ダッシュボード使用方法
：https://github.com/kubernetes/dashboard
## using

**動作確認**

- kubernatesの場合ポートを繋ぐ処理が必要
```
kubectl get pod // pod名確認
kubectl port-forward <pod名> 3000:3000　// 接続
```
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
- kubernatesの場合環境のクリーンアップ

```
kubectl get services // serviceの確認
kubectl delete services <サービス名> // serviceの削除
kubectl get deployment // deploymentの確認
kubectl delete deployment <deployment名>　// deploymentの削除
```
[GinとAPIに関する参考資料](https://go.dev/doc/tutorial/web-service-gin)

[curlに関する参考資料](https://qiita.com/yasuhiroki/items/a569d3371a66e365316f)

[環境構築関する参考資料](https://www.utakata.work/entry//golang/tutorial/2-docker-go-gin)

[kubernates参考資料](https://kubernetes.io/ja/docs/tasks/access-application-cluster/service-access-application-cluster/)

[kubernatesクリーアップ](https://qiita.com/superbrothers/items/0dca5d2a10727fc14734)

[kubernates全体参考資料](https://techblog.istyle.co.jp/archives/6296)
