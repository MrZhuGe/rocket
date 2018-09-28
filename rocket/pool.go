package rocket

import (
	"math"
	"sync/atomic"
)

const (
	DefaultPoolCap = math.MaxInt32
)

type task func() 

type Pool struct {

	// 当前正在执行的goroutine数量
	running int32

	// 待执行任务队列
	tasks *List

	// 任务添加
	taskChan chan struct{}

	// 工作队列的最大数量
	workerChan chan struct{}

	closed int32
}

func NewPool() (*Pool, error) {
	return NewPoolCap(0)
}

func NewPoolCap(cap int32) (*Pool, error) {

	p := &Pool{
		tasks:      NewList(),
		taskChan:   make(chan struct{}, DefaultPoolCap),
		workerChan: make(chan struct{}, cap),
	}

	return p, nil
}

// 加入执行任务队列
func (p *Pool) Add(t task) {
	p.tasks.PushBack(t)
	p.taskChan <- struct{}{}

	// 判断是否创建新的worker
	if p.Tasks() > 1 || p.Running() == 0 {
		p.Worker()
	}
}

// 获取空闲任务队列
func (p *Pool) Worker() {

	if cap(p.workerChan) > 0 {
		if int(p.Running()) == cap(p.workerChan) {
			return
		}
		p.workerChan <- struct{}{}
	}
	p.incRunning()

	go func() {
		for !p.GetClosed() {
			select {
			case <-p.taskChan:
				if t := p.tasks.PopFront(); t != nil {
					t.(task)()
				} else {
					goto done
				}
			default:
				goto done
			}
		}

	done:
		p.decRunning()
		if cap(p.workerChan) > 0 {
			<-p.workerChan
		}
	}()
}

// 查询当前等待处理的任务总数
func (p *Pool) Tasks() int {
	return p.tasks.Len()
}

// 正在执行任务队列
func (p *Pool) Running() int32 {
	return atomic.LoadInt32(&p.running)
}

func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

// 设置工作队列的开启状态
func (p *Pool) SetClosed(v bool) {
	if v {
		atomic.StoreInt32(&p.closed, 1)
	} else {
		atomic.StoreInt32(&p.closed, 0)
	}
}

// 获得工作队列的开启状态
func (p *Pool) GetClosed() bool {
	return atomic.LoadInt32(&p.closed) > 0
}