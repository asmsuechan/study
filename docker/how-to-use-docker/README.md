# Docker

とりあえず
```
$ brew install docker
$ brew install docker-machine
```

## dinghyの役割
> OSXのDocker MachineのVolumes機能をNFSに入れ替えてくれる

dinghyはvagrantでdocker入りのマシンを立ち上げたりファイル共有をNFSにしてくれる

> docker-machineコマンドの代わりにdinghyコマンドを使って操作

とあるのでdocker-machineに何らかの不便点があったからdinghyへの移行が進んだと思われる

### dinghyの導入
```
$ brew tap codekitchen/dinghy
$ brew install dinghy
$ dinghy create --provider virtualbox
```

```
$ dinghy help create
```

でcreateコマンドのオプションが見られる

ちょっとエラー起きた
http://qiita.com/niiyz/items/70580164551c710a75d3

## quay.ioとは

## Docker
各種ソフトウェアのコンテナ

## docker-compose
コンテナをまとめる

# Dockerを動かす

こんなエラーが出たらmacを再起動
```
$ docker ps
Cannot connect to the Docker daemon. Is the docker daemon running on this host?
```

```
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
$ docker ps #process idを調べる
$ docker exec -it 39bda6340dd0 /bin/bash
```

ここがいい感じにまとまっている
http://qiita.com/7kaji/items/e4ad05188df6d05aff97

docker-composeを使う


