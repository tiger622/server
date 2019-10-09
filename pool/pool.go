package pool

import (
	"sync"
	"time"
)

type Pool struct {
	// 工作池最大数
	poolsize int16
	// 工作总数(不超过最大数)
	allworker int16
	// 空闲工作队列
	workers []*worker
	// 工作池检查间隔
	expiryDuration time.Duration
	lock           sync.Mutex
}
