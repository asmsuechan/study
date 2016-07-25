# MacにROSをインストール
**結論：失敗**

## dockerでrosをインストール
dockerとdocker-machineがインストールされていることが前提になります。

あとは[公式サイト](http://wiki.ros.org/docker/Tutorials/Docker)を参考にしてインストールしました。

### こんなエラーが出た時は
```
$ docker pull ros
Using default tag: latest
Post http:///var/run/docker.sock/v1.20/images/create?fromImage=ros%3Alatest: dial unix /var/run/docker.sock: no such file or directory.
* Are you trying to connect to a TLS-enabled daemon without TLS?
* Is your docker daemon up and running?
```

これで解決
```
$ docker-machine start default
$ eval "$(docker-machine env default)"
```
このevalは起動する各シェルで実行するか適当にシェルの設定ファイルに突っ込むかしてください。

以下のstackoverflowにありました。
http://stackoverflow.com/questions/32425717/is-your-docker-daemon-up-and-running-problems-with-docker-hello-world-tutoria

## dockerからrvizを起動
ここからは[こちらのサイト](http://tech-sketch.jp/2015/10/ros-docker-1.html)を参考にしました。

中段の「3. Dockerコンテナ内でGUIツール(Rviz)を使おう」からを実行していきます。

参考サイトのコマンドを実行するとエラーを吐きます。
```
 docker run -it --rm \
 >     --name rviz \
 >     --env ROS_HOSTNAME=rviz \
 >     --env ROS_MASTER_URI=http://master:11311 \
 >     --env DISPLAY=unix$DISPLAY \
 >     --volume="/tmp/.X11-unix:/tmp/.X11-unix:rw" \
 >     osrf/ros:indigo-desktop-full \
 >     rosrun rviz rviz
 rviz: cannot connect to X server unix/private/tmp/com.apple.launchd.IUheFkztTD/org.macosforge.xquartz:0
 ``

XQuartzがエラー吐いてますね。あとrviz立ち上げるだけなのですがここが一番根が深そうです。

参考サイトはUbuntuでdocer立ててros動かしているのでうまくMacで動くようにしなければいけません。

自分はここでギブアップしました。MacでdockerのプログラムのGUIを実行する方法を学ばなければいけません。
