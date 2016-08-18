# [mysqlが起動しない問題]

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

その1
```
2016-07-29 10:10:42 974 [Note] IPv6 is available.
2016-07-29 10:10:42 974 [Note]   - '::' resolves to '::';
2016-07-29 10:10:42 974 [Note] Server socket created on IP: '::'.
2016-07-29 10:10:42 974 [ERROR] /usr/local/Cellar/mysql56/5.6.31/bin/mysqld: Can't find file: './mysql/user.frm' (errno: 13 - Permission denied)
2016-07-29 10:10:42 974 [ERROR] Fatal error: Can't open and lock privilege tables: Can't find file: './mysql/user.frm' (errno: 13 - Permission denied)
160729 10:10:42 mysqld_safe mysqld from pid file /usr/local/var/mysql/ats.local.pid ended
ats:mysql
```

その2
```
160817 22:04:34 mysqld_safe Starting mysqld daemon with databases from /usr/local/var/mysql
2016-08-17 22:04:35 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
2016-08-17 22:04:35 0 [Note] /usr/local/Cellar/mysql56/5.6.31/bin/mysqld (mysqld 5.6.31) starting as process 55724 ...
2016-08-17 22:04:35 55724 [Warning] Setting lower_case_table_names=2 because file system for /usr/local/var/mysql/ is case insensitive
2016-08-17 22:04:35 55724 [Note] Plugin 'FEDERATED' is disabled.
/usr/local/Cellar/mysql56/5.6.31/bin/mysqld: Unknown storage engine 'InnoDB'
2016-08-17 22:04:35 55724 [ERROR] Can't open the mysql.plugin table. Please run mysql_upgrade to create it.
2016-08-17 22:04:35 55724 [Note] InnoDB: Using atomics to ref count buffer pool pages
2016-08-17 22:04:35 55724 [Note] InnoDB: The InnoDB memory heap is disabled
2016-08-17 22:04:35 55724 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
(中略)
InnoDB: to create the InnoDB data files, but log file creation failed.
InnoDB: If that is the case, please refer to
InnoDB: http://dev.mysql.com/doc/refman/5.6/en/error-creating-innodb.html
2016-08-17 22:21:15 9454 [ERROR] Plugin 'InnoDB' init function returned error.
2016-08-17 22:21:15 9454 [ERROR] Plugin 'InnoDB' registration as a STORAGE ENGINE failed.
2016-08-17 22:21:15 9454 [ERROR] Unknown/unsupported storage engine: InnoDB
2016-08-17 22:21:15 9454 [ERROR] Aborting
```

ログの更新をリアルタイムで覗くには別タブ開いてtailfを実行してmysqlを起動します。
```
$ tail -f /usr/local/var/mysql/*.local.err
# 別タブ
$ mysql.server start
```

### chown _mysqlで治らない場合
いろんな場所にsudo chown -R _mysql:_mysql /usr/local/var/mysql/とか書いてあるのですが、結構治りません。

そういう場合はとりあえず再インストール、バージョンを変える、等しましょう。

## 参考
http://qiita.com/shogo807/items/f90269818040b2b2781f
