package rocket

import (
	"container/list"
	"sync"
)

// 并发安全的双向链表
type List struct {
	mu *sync.RWMutex
	list *list.List
}

func NewList() *List {
	return &List{
		mu: new(sync.RWMutex),
		list:list.New(),
	}
}

// 往链表头入栈数据项
func (l *List) PushFront(v interface{}) *list.Element {
	l.mu.Lock()
	e := l.list.PushFront(v)
	l.mu.Unlock()
	return e
}

// 往链表尾入栈数据项
func (l *List) PushBack(v interface{}) *list.Element {
	l.mu.Lock()
	r := l.list.PushBack(v)
	l.mu.Unlock()
	return r
}

// 从链表尾端移除数据项
func (l *List) PopBack() interface{} {
	l.mu.Lock()
	if e := l.list.Back(); e != nil {
		v := l.list.Remove(e)
		l.mu.Unlock()
		return v
	}
	l.mu.Unlock()
	return nil
}

// 从链表头端移除数据项
func (l *List) PopFront() interface{} {
	l.mu.Lock()
	if e := l.list.Front(); e != nil {
		v := l.list.Remove(e)
		l.mu.Unlock()
		return v
	}
	l.mu.Unlock()
	return nil
}

// 移除数据项
func (l *List) Remove(e *list.Element) interface{} {
	l.mu.Lock()
	r := l.list.Remove(e)
	l.mu.Unlock()
	return r
}

// 获取链表头值
func (l *List) Front() interface{} {
	l.mu.RLock()
	if e := l.list.Front(); e != nil {
		l.mu.RUnlock()
		return e.Value
	}

	l.mu.RUnlock()
	return nil
}

// 获取链表尾值
func (l *List) Back() interface{} {
	l.mu.RLock()
	if e := l.list.Back(); e != nil {
		l.mu.RUnlock()
		return e.Value
	}

	l.mu.RUnlock()
	return nil
}

// 获取链表长度
func (l *List) Len() int {
	l.mu.RLock()
	length := l.list.Len()
	l.mu.RUnlock()
	return length
}