# jsのテストについて
自分が作っている[ライブラリ](https://github.com/asmsuechan/jquery_push_notification)にもテストはあるのですがjsのテストは初めてだったのでメモします。

## karma+jasmine
作ったライブラリがview用だったのでkarmaとjasmineでテストをしました。

karmaは指定したブラウザでテストを実行してくれるツールで、jasmineはテストフレームワークです。

### karma
まずはインストールします。
```
$ npm install -g karma
$ karma init
```
karma initすると対話形式で設定画面が出てきます。ここで設定ファイル(karma.conf.js)を作成します。

フレームワークはjasmineを使い、ブラウザは`ChromeとSafariとFirefox`を選択しました。

設定は後からでも変更可能なので全部エンターキー押してしまっても良さそうです。

自分の最終的なkarma.conf.jsはこうなっています  
https://github.com/asmsuechan/jquery_push_notification/blob/master/karma.conf.js

### jasmine
まずはインストールします。
```
$ npm install -g jasmine
```

次にbowerでいろいろインストールします

```
$ bower init
$ bower install --save-dev jquery
$ bower install --save-dev jasmine-jquery#~1
```

インストールしたらkarma.conf.jsのfilesに追記していきます
```
files: [
  'bower_components/jquery/dist/jquery.min.js',
  'bower_components/jasmine-jquery/lib/jasmine-jquery.js'
  ]
```

### テストコマンドをnpmのタスクに追加

package.jsonに以下を追加します
```
"scripts": {
  "test": "node_modules/karma/bin/karma start"
},
```
これで`npm test`でテストが走ります。

とりあえず実行したら指定したブラウザが開きます。これがテスト画面です。

### テストコードを書く
テストコードを置くディレクトリを作成してkarma.conf.jsのfilesに追記します。

```
files: [
  'test/unit/*Spec.js'
  ]
```
テストコードはtest/unit/に置き、*Spec.jsという命名規則にします。

ここはよしなに書きます。

# 参考
ありがてえありがてえ

http://qiita.com/shinofara/items/b3677ffdfc0c7e45e8d4

https://tech.recruit-mp.co.jp/front-end/post-5299/

http://qiita.com/ngyuki/items/d36e500b55e1f627acab
