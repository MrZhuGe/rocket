package rocket

// 默认协程池
var defaultPool, _ = NewPool()

func Add(t task) {
	defaultPool.Add(t)
}

func Running() int32 {
	return defaultPool.Running()
}

func SetClosed(v bool) {
	defaultPool.SetClosed(v)
}