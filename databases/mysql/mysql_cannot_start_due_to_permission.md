# mysqlが起動しない問題

ほとんどの場合設定の権限がおかしくなっています。

### いつまでたっても起動しない
```
$ mysql.server start
Starting MySQL
.....................................................................^C
```

既にmysqlが立ってる場合です。 `mysql -uroot` で入ってみましょう。

### 権限エラー
一番多くて厄介なやつです。こんなやつです。権限エラーです。

```
$ mysql.server start
Starting MySQL
. ERROR! The server quit without updating PID file (/usr/local/var/mysql/user/pid)
```

mysqlの起動ユーザーと設定ファイル(/usr/local/var/mysql/*)のユーザーを同じにしなくてはいけません。

mysqlの起動ユーザーを確認するために設定ファイルを見ます。
```
$ mysql --help | grep my.cnf
```

既存プロセスがないかチェックします。いたらmysqlと名のつくプロセス全部殺しマンに殺してもらいます。

```
$ ps aux | grep mysql
$ kill `ps -ef | grep mysql | awk '{print $2;}'`
```

ここでもう一度mysqlを起動してみます
```
$ mysql.sersver start
```

### ログを見よう
```
ls -la /usr/local/var/mysql
```
ここにある*.local.errがエラーログファイルです。

```
$ cat /usr/local/var/mysql/*.local.err
```

ここの下の方が最新のログです。こんな感じになってます。

```
2016-07-29 10:10:42 974 [Note] IPv6 is available.
2016-07-29 10:10:42 974 [Note]   - '::' resolves to '::';
2016-07-29 10:10:42 974 [Note] Server socket created on IP: '::'.
2016-07-29 10:10:42 974 [ERROR] /usr/local/Cellar/mysql56/5.6.31/bin/mysqld: Can't find file: './mysql/user.frm' (errno: 13 - Permission denied)
2016-07-29 10:10:42 974 [ERROR] Fatal error: Can't open and lock privilege tables: Can't find file: './mysql/user.frm' (errno: 13 - Permission denied)
160729 10:10:42 mysqld_safe mysqld from pid file /usr/local/var/mysql/ats.local.pid ended
ats:mysql
```

ログの更新をリアルタイムで覗くには別タブ開いてtailfを実行してmysqlを起動します。
```
$ tail -f /usr/local/var/mysql/*.local.err
# 別タブ
$ mysql.server start
```

## 参考
http://qiita.com/shogo807/items/f90269818040b2b2781f
