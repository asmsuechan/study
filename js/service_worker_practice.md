# [ServiceWorkerについて]
## ServiceWorkerとは
ServiceWorkerはブラウザのバックグラウンドに常駐するプログラムのようなものです。オフラインでも動作するうえにページを開いていない状態でもユーザーへアクションを起こすことができます。

↓開発ツールから見えるServiceWorker
![開発ツール](https://github.com/asmsuechan/study/raw/master/assets/sw_fb.png "Chrome開発ツールのApplication")

また、ServiceWorkerはほとんど未来の技術と思われるほど新しい技術なのでだいたいのドキュメントには

> **ほとんどの機能は正しく動作しません。**

のようなことが書かれています。規格はあっても実装段階までいっていません。実際まともに動くのはChromeとFirefoxくらいです。

Chromeのバージョン対応については下で触れています。

## できること
Clientとして判断されたページのイベントを受け取ることができます。

受け取れるイベントは

* fetch
* push
* activate
* install

です。
## プッシュ通知
ServiceWorkerにpushイベントハンドラを追加することによりいい感じにプッシュ通知を出すことができます。(プッシュ通知を出すだけならpushイベント関係なく出せます)

## キャッシュ
## ブラウザ対応
ServiceworkerはChrome40~

showNotification()はChrome42~

postMessage()はChrome45~

## 参考
[ServiceWorker API(MDN)](https://developer.mozilla.org/ja/docs/Web/API/ServiceWorker_API)
