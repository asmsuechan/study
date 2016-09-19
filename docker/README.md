# [Dockerを使う]

## 必要なもの
* docker
* docker-compose
* dinghy
* virtualbox

virtualbox以外はbrewで入ります。

## docker
dockerは仮想化技術の一つで、システムはホストOSを使いその上にコンテナを乗せていてプロセスを実行します。

## docker-compose
docker-composeは複数のコンテナをyamlで管理するツールです。

特別なことはしておらずdocker run等で指定できるオプションをそのままyamlにしただけだが見通しよく管理することができ保守性が高いです。

例えばmysqlは以下のように書けます。
```
# mysql
mysql:
  image: mysql
  volumes_from:
    - data-mysql
  ports:
    - "3306:3306"
  environment:
    MYSQL_ROOT_PASSWORD: password
```

`docker-compose run`を頻繁に使ってます。
[ここ](http://qiita.com/taka4sato/items/f03004e449538b325b5e#docker-compose-run)が参考になります。

**参考**
http://qiita.com/y_hokkey/items/d51e69c6ff4015e85fce

## dinghy
docker-machineより高速にしたものです。

[この](http://qiita.com/masuidrive/items/d71ee1881fffb6ad098f)参考サイトによると18倍くらい早くなるみたいです。

> VMWareで30秒、Virtualboxで90秒かかる中規模のRailsアプリの起動が5秒に短縮されます
