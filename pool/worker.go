package pool

// 工作执行函数
type work interface {
	Execute() error
}

// 任务的执行者,从channel中取任务(每个worker是一个goroutine)
type worker struct {
	task chan work
	pool *Pool
}
