# go-next-monorepo

一個網頁伺服器範本，把靜態前端（預設是 nextjs）打包後，跟後端（golang, gin）一起打包成一個可執行檔

# Quick Start

```
$ make doctor
$ make dep
$ make dev
$ make build
```

# feature

-   輸出只有一個靜態編譯的可執行檔，可以到處帶著走
-   開發模式下，前後端都有 hotreload（前端是 nextjs 自帶、後端是用 nodemon 監看）
-   超小的 docker image size（感謝 Golang 的靜態編譯）
-   websocket（可以移除）
-   幫你整合好 api 路由和 static files 的路由了
-   完善的 Makefile，包含相依性安裝、開發模式、打包全部都在裡面（詳情請看 `make help`）
-   高度可客製化，修改 `backend/build.sh` 和 `frontend/build.sh` 就可以更改前後端編譯配置

## frontend

> [!NOTE]  
> 前端使用的包管理器是 [pnpm](https://pnpm.io)，如果你不喜歡，請自行修改 `Makefile` 和 `frontend/build.sh`，但我還是強烈推薦 pnpm，他真的很棒！

前端雖然預設是 nextjs，但是你只要修改 `frontend/build.sh`，把最終的靜態檔案放到 `frontend/out`，接著理論上 `make build` 時就會幫你把檔案包進去了

> [!WARNING]  
> 我沒實驗過，如果有人成功歡迎回報

## backend

後端的資料夾比較複雜一點，首先，先來看 main.go

### main.go

main.go 定義了一個函式 `run(addr string) error`，裡面會用 gin 開啟一個 http server，並依照開發/發布模式設定好靜態檔案的路由（下面會說明）。正常情況下你不需要動到 main.go，不過有一個情況例外：移除 websocket 支援

### 移除 websocket 支援

如果你很確定不需要 websocket 支援，或是你覺得我寫得很爛（確實，但我懶得改了，我本來想弄成 socket.io 那樣，但是功力不夠 QQ）。可以移除以下幾行/檔案

1. 移除 `backend/internal/websocket/` 整個目錄
2. 移除 `backend/main.go` 中的

```
	"backend/internal/websocket"
```

和

```
	io := websocket.Route(r)
```

，並且修改

```
	api.Route(r, io)
```

成

```
	api.Route(r)
```

3. 修改 `backend/api/api.go` 中所有關於 `io` 的部份（這裡我懶得打了 ==，我相信你會）

### 新增 API

後端就是拿來放 API 的啦！按照我的設計，新增 API 全部都是放在 `backend/api/` 目錄下，你可按照 API endpoint 再去細分 `backend/api/user/`、`backend/api/post/` 之類的，總之，看你開心～

### 編譯變數

`backend/main.go` 中定義了四個編譯變數，他們會在 `backend/build.sh` 中塞值進去，分別是

-   Mode: `"debug"` 或是 `"release"`
-   Version：git tag 的版本
-   CommitHash：執行 `backend/build.sh` 時的 git commit hash
-   BuildTime：執行當下時間

你可以執行編譯完的可執行擋 `./main -v` 看這些訊息

# TODO

-   add Action to build automatically
-   ~~docker image ?~~
-   add descripton of make command
