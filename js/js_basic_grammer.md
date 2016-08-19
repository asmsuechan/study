# [jsの基礎的な文法メモ]
js弱者なので基礎の基礎から勉強します

### コンストラクタ
jsにはクラスの代わりにコンストラクタからインスタンスを作成します。


参考

[JavaScriptのクラス？コンストラクタ？？](http://qiita.com/takeharu/items/010752b1427773558f7c)

### Promise
非同期処理の後に何がしかの処理をしたい時に使います。

#### thenメソッド
promiseの処理が終了した後に実行したい処理を書きます。

```
//getMessage()で新着メッセージをfetchして取ってくる
getMessage().then(function(){
  self.registration.showNotification(title, {
    body: body,
    icon: icon,
    tag: tag
  })
);
```

thenの終わりにセミコロンがあるとエラーが出ます。

#### waitUntilメソッド
このメソッド内の処理が終わるまでサービスワーカーはインストールされません。

#### Promise参考
[JavaScript Promiseの本](http://azu.github.io/promises-book/)
[JavaScript プログラミング講座 Promiseについて](http://hakuhin.jp/js/promise.html)
[JS.next Promiseについて](http://js-next.hatenablog.com/entry/2013/11/28/093230)
