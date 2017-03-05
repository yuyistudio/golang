// 实现并发的非阻塞式缓存
// http://docs.plhwin.com/gopl-zh/ch9/ch9-07.html

package main

import (
	"fmt"
	"sync"
)

type result struct {
	value string
	err   error
}
type Func func(string) (string, error)

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 第一次尝试获取缓存时
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		// 以后每次获取缓存都直接走到这里
		go e.deliver(req.response) // Get调用者在response上等待
	}
}

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready) // 关闭ready,让deliver函数可以走下去(此时该缓存的获取者应该都卡在ready上等着)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready // 等待第一次缓存完成(call函数完成)
	// Send the result to the client.
	response <- e.res // 返回结果给调用者
}

func addPostfix(prefix string) (string, error) {
	return prefix + "|post", nil
}

func main() {
	m := New(addPostfix)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			fmt.Println(m.Get(fmt.Sprint(i)))
			wg.Done()
		}()
	}
	wg.Wait()
}
