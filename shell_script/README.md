# シェルスクリプト勉強記録

### ファイルの存在判定

ファイルが存在する
```
if [ -f ~/.vimrc ];then
  echo "~/.vimrcは存在します"
fi
```

ファイルが存在しない
```
if [ ! -f ~/.vimrc ];then
  echo "~/.vimrcは存在しません"
fi
```

### ディレクトリの存在判定

ディレクトリが存在する
```
if [ -d ~/dotfies ];then
  echo "~/dotfilesディレクトリは存在します"
fi
```

ディレクトリが存在しない
```
if [ ! -d ~/dotfies ];then
  echo "~/dotfilesディレクトリは存在しません"
fi
```

### コマンドの存在判定

正規表現でチェックします

コマンドが存在しない場合
```
if [[ `type brew` =~ (?!found)$ ]];then
  echo "brew not exists"
fi
```
コマンドが存在しない場合`brew not found`と出るので正規表現で「末尾にfoundがあるか」のマッチングをしています。

## 文法
### バッククオーテーション
バッククオーテーション(\`)でくくると内部でコマンドを実行して返ってきた値を貼り付けます。例えば
```
echo `ls`
```
でlsコマンドで返ってきた文字列を出力します。これはif文でよく使うと思います。

### if文の[ ]
[ ]はtestコマンドのエイリアスです。

testコマンドは評価値に応じて0か1を返すコマンドです。

## 参考
ありがてえありがてえ
http://qiita.com/toshihirock/items/461da0f60f975f6acb10

http://shellscript.sunone.me/if_and_test.html
