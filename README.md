# echo-sample

### 導入
- ライブラリをインストール
```
$ go get -u github.com/labstack/echo/...
```

- サンプル
```
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hoge", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

- サーバー起動
```
$ go run main.go
```

- アクセス
http://localhost:1323/hoge