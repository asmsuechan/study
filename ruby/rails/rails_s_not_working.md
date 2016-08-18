# [rails_sが動かない]
こんなエラーが出ました。
```
$ rails s
bin/rails:6: warning: already initialized constant APP_PATH
/Users/ryouta/soccker_channel/soccer_channel/bin/rails:6: warning: previous definition of APP_PATH was here
Usage: rails COMMAND [ARGS]

The most common rails commands are:
 generate    Generate new code (short-cut alias: "g")
 console     Start the Rails console (short-cut alias: "c")
 server      Start the Rails server (short-cut alias: "s")
 dbconsole   Start a console for the database specified in config/database.yml
             (short-cut alias: "db")
 new         Create a new Rails application. "rails new my_app" creates a
             new application called MyApp in "./my_app"

In addition to those, there are:
 destroy      Undo code generated with "generate" (short-cut alias: "d")
 plugin new   Generates skeleton for developing a Rails plugin
 runner       Run a piece of code in the application environment (short-cut alias: "r")

All commands can be run with -h (or --help) for more information.
```

どうやら/usr/local/mysql/lib/libmysqlclient.18.dylibがないせいらしいです。

が、 `sudo find / | grep libmysqlclient.18.dylib` しても見つかりませんでした。

原因はまたmysql5.7でした。5.6にダウングレードします。

```
$ brew uninstall mysql
$ brew install homebrew/versions/mysql56
```

さあ今度はmysqlが起動しないぞ！
[mysqlが起動しない問題](https://github.com/asmsuechan/study/blob/master//databases/mysql/mysql_cannot_start_due_to_permission.md)

にっちもさっちも行かなかったのでダメ元でもっかい再インストール
```
$ brew uninstall mysql
$ brew install mysal56
```

なんとこれで通りました。。。

そしてこれでrails sも通りました。
