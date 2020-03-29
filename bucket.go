package bucket

import (
	"sync"
	"time"
)

func NewBucket(duration time.Duration) *Bucket {

	b := new(Bucket)
	b.mu = new(sync.Mutex)
	b.duration = duration
	b.first = true
	b.change = false
	b.setNext()

	return b
}

type Bucket struct {
	duration time.Duration
	messages []Message
	mu       *sync.Mutex
	next     int64
	change   bool
	first    bool
}

func (p *Bucket) now() int64 {
	return time.Now().UnixNano()
}

func (p *Bucket) setNext() {
	p.next = p.now() + int64(p.duration)
}

func (p *Bucket) call(callback func(messages []Message)) {

	p.mu.Lock()
	messages := make([]Message, len(p.messages))
	copy(messages, p.messages)
	p.mu.Unlock()

	callback(messages)
	p.messages = make([]Message, 0)
	p.setNext()
}

func (p *Bucket) Push(message Message) {
	p.mu.Lock()
	if len(p.messages) == 0 {
		p.change = true
	}else {
		p.change = false
	}
	p.messages = append(p.messages, message)
	p.mu.Unlock()
}

func (p *Bucket) Pop(callback func(messages []Message)) {
	for {
		if p.now() >= p.next {
			p.call(callback)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func (p *Bucket) First(callback func(messages []Message)) {
	for {
		if p.change && p.first {
			p.first = false
			p.call(callback)
		} else if p.now() >= p.next {
			p.call(callback)
			p.first = true
		}
		time.Sleep(300 * time.Millisecond)
	}
}
