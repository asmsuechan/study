# [fetchAPIにパラメーターを乗せる]
こんな感じで乗せます。

ちゃんとコンバートしなきゃいけないみたいですね。

```
fetch('/user', {
  method: 'POST',
  body: convertObjectToFormData({ registration_id: registration_id}),
  credentials: 'include'
})

function convertObjectToFormData(query) {
    var formData = new FormData();
    for (var key in query) {
        formData.append(key, query[key]);
    }
    return formData;
}
```

## 参考
http://stackoverflow.com/questions/34692554/url-query-string-in-fetch-api-in-javascript
