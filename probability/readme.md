## 機率

### 使用方式

```shell
go get -u github.com/yuyuancha/probability
```

創建機率權重，依照權重給定索引。

```
w := probability.CreateWeight(1, 30, 20, 10, 1, 1, 1)
fmt.Println(w.GetCount()) // 打印總共的數量
for i := 0; i < 5; i++ {
	fmt.Println(w.GetIndexByWeight()) // 按照權重給定索引
}
```
