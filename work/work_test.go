package work

import (
	"log"
	"testing"
	"sync"
)

// names提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	//time.Sleep(time.Second)
}

func TestWork(t *testing.T) {
	// 使用1000个goroutine来创建工作池
	p := New(1000)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行。当Run返回时
				// 我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，等待所有现有的
	// 工作完成
	p.Shutdown()
}
