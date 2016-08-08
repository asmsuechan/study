# [after_commitはdestroy時にも走る]
当たり前の話ですが、after_commitはすべてのcommit後に走るのでもちろんdestroyの時も走ります。

ここが考慮漏れしており以下のエラーが出ました。

`RuntimeError: Can't modify frozen hash`

これはdestroyで削除済みのhashに対してafter_commitで操作しようとしたために起こりました。

以下のように対処しました。

```
after_commit :set_user_id, on: :create
```

on: :createをつけることによりcreate時のみしか動かなくしました。これで解決です。

### 参考
http://apidock.com/rails/ActiveRecord/Transactions/ClassMethods/after_commit
