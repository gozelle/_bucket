package bucket

import (
	"fmt"
	"testing"
	"time"
)

type M string

func (m M) String() string {
	return string(m)
}

func TestNewBucket(t *testing.T) {
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
}
