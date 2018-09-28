# rocket
Easy to use Golang goroutine pool,performance and elegance coexist

rocket是一个高性能的协程池，实现了对大规模goroutine的调度管理，尽可能的复用了已有的goroutine，减少了goroutine的频繁创建销毁带来的系统资源浪费问题，并且允许开发者限制goroutine的数量，达到更好的效果。  
## 功能
1.自动调度goroutine，复用goroutine  
2.资源复用，在高并发下，节省大量资源和内存占用  
3.增加任务方便简单  
