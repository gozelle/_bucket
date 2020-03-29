# Bucket

计时触发消息

示例：

```go
	bucket := NewBucket(2 * time.Second)
	go func() {
		for {
			bucket.Push(M(time.Now().String()))
			time.Sleep(600 * time.Millisecond)
		}
	}()

	bucket.First(func(messages []Message) {
		fmt.Println(len(messages), messages)
	})
```