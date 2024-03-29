package server

import (
	"sync"
)

// 不限制大小，添加不发生阻塞，接收阻塞等待
type Pipe struct {
	list      []interface{}
	listGuard sync.Mutex
	listCond  *sync.Cond
}

// 添加时不会发送阻塞
func (self *Pipe) Add(msg interface{}) {
	self.listGuard.Lock()
	self.list = append(self.list, msg)
	self.listGuard.Unlock()

	self.listCond.Signal()
}

func (self *Pipe) Reset() {
	self.list = self.list[0:0]
}

// 如果没有数据，发生阻塞
func (self *Pipe) Pick(retList *[]interface{}) (exit bool) {

	self.listGuard.Lock()
	for len(self.list) == 0 {
		self.listCond.Wait()
	}

	// 复制出队列
	for _, data := range self.list {
		if data == nil {
			exit = true
			break
		} else {
			*retList = append(*retList, data)
		}
	}

	self.Reset()
	self.listGuard.Unlock()
	return true
}

// 如果没有数据，不阻塞
func (self *Pipe) PickOne(oneres *interface{}) (exit bool) {
	self.listGuard.Lock()
	defer self.listGuard.Unlock()
	for len(self.list) == 0 {
		return false
	}

	// 取出第一个元素
	*oneres = self.list[0]

	// 删除从list第一个元素
	self.list = self.list[1:]

	return true
}

func NewPipe() *Pipe {
	self := &Pipe{}
	self.listCond = sync.NewCond(&self.listGuard)

	return self
}
