// runner包管理处理任务的运行和生命周期
package runner

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

// Runner在给定的超时时间内执行一组任务，
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统
	// 发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	timeout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的 26 // 函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超市时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接受到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New return a can use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将一个任务附加到Runner上。这个任务是一个
// 接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
		// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已注册的任务
		task(id)
	}

	return nil
}

// 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}
