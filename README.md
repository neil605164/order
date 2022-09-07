# 注意事項

## 安裝

### Docker

- 安裝教學　[連結](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04)

### Docker-compose

- 安裝教學 [連結](https://www.digitalocean.com/community/tutorials/how-to-install-docker-compose-on-ubuntu-16-04)

### Golang

1. 點此 [連結](https://golang.org/dl/) 取得 golang 新版本

2. 解壓縮檔案
   
```
$ sudo tar -C /usr/local -xzf ~/Downloads/go1.14.linux-amd64.tar.gz
```

3. 確認 __PATH__ 有加上 `/usr/local/go/bin`
   
```
$ echo $PATH | grep "/usr/local/go/bin"
```

### Swag
安裝 swagger 文件　[連結](https://github.com/swaggo/swag)

```
$ go get -u github.com/swaggo/swag/cmd/swag
```

---
## 啟動服務
執行腳本 `RunService.sh`

```
$ sh RunService.sh
```

---
## 建立 db table & 關聯
```
ENV=local go run migrations/init.go
```

## 新增假資料
```
INSERT INTO `products` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, '產品名稱一', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
```

---
## DB 結構關聯圖
- 文件 [連結](https://app.diagrams.net/#G1zoB9UmNjqfcUc-vgMqUYULoSbZvaRm9t)

## api 文件連結

**注意** 必須啟動服務後，連結才可以開啟
- 文件 [連結](http://localhost:9999/api/v1/swagger/index.html#/order/post_order)