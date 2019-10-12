package work

import "sync"

// work包的目的是展示如何使用无缓冲的通道来创建一个 goroutine 池，
// 这些 goroutine 执行 并控制一组工作，让其并发执行。
// Worker必须满足接口类型才能使用工作池
type Worker interface {
	Task()
}

// Pool提供一个goroutine池，这个池可以完成
// 任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
