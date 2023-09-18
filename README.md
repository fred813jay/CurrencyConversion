# 簡介

這是一個使用Gin框架提供貨幣轉換API的Go程式。
它將一個貨幣金額從一種貨幣轉換為另一種貨幣，將結果四捨五入到小數點後兩位，並使用逗號進行千位分隔的格式化。

## 程式版本與使用框架

- Go 1.20.2
- Gin 1.9.1

### 使用說明

- 使用 go test 命令執行單元測試。

- 使用 go run main.go 命令執行程式或是 go build 打包後執行程式。

- 查詢參數：<br>
   1.source：原始貨幣 (TWD/JPY/USD)。<br>
   2.target：目標貨幣 (TWD/JPY/USD)。<br>
   3.amount：要轉換的金額(格式:$金額，使用逗號進行千位分隔)。

- 輸入範例:?source=USD&target=JPY&amount=$1,525

- 程式執行後，使用此連結[查看輸出結果](http://localhost:8080/?source=USD&target=JPY&amount=$1,525)
