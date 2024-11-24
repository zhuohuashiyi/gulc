package gulc


import (
	"sync"
	"time"
	"context"
)


type Future struct {
	ch chan interface{}
}

func (f *Future) Get() interface{} {
	var res interface{}
	res, ok := <- f.ch
	for !ok {
		time.Sleep(1 * time.Millisecond)
		res, ok = <- f.ch
	}
	return res
}

type Work struct {
	work func(args []interface{}) interface{}
	args []interface{}
	res  *Future
}


type Worker struct {
	taskChannel chan Work
	isCore      bool
	keepAliveTime  time.Duration
	wg *sync.WaitGroup
}

func NewWorker(taskChannle chan Work, isCore bool, keepAliveTime time.Duration, wg *sync.WaitGroup) *Worker {
	return &Worker{taskChannel: taskChannle, isCore: isCore, keepAliveTime: keepAliveTime, wg: wg}
}


func (w *Worker) Work(ctx context.Context) {
	defer w.wg.Done()
	timer := time.NewTimer(w.keepAliveTime)
	for {
		select {
		case <- timer.C:
			if !w.isCore {
				return
			}
		case <- ctx.Done(): 
			return
		case task := <- w.taskChannel:
			res := task.work(task.args)
			task.res.ch <- res
			timer.Reset(w.keepAliveTime)
			continue
		}
	}
}


type ThreadPool struct {
	workers  []*Worker
	coreThreads int
	maxThreads int
	keepAliveTime time.Duration
	taskChannel chan Work
	ctx context.Context
	cancel context.CancelFunc
	threshold int
	discount int
	wg *sync.WaitGroup
}

func NewThreadPool(coreThreads, maxThreads int, keepAliveTime time.Duration, threshold, discount int) *ThreadPool {
	taskChannel := make(chan Work, maxThreads * 2 + 1)
	workers := make([]*Worker, 0)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < coreThreads; i++ {
		worker := NewWorker(taskChannel, true, keepAliveTime, &wg)
		go worker.Work(ctx)
		workers = append(workers, worker)
		wg.Add(1)
	}
	t := &ThreadPool{workers: workers, coreThreads: coreThreads, maxThreads: maxThreads, keepAliveTime: keepAliveTime, taskChannel: taskChannel, ctx: ctx, cancel: cancel, threshold: threshold, discount: discount, wg: &wg}
	go t.daemon()
	return t
}

// Submit 往协程池中提交任务
func (t *ThreadPool) Submit(function func(args []interface{}) interface{}, args []interface{}) *Future {
	resChannel := make(chan interface{}, 1)
	future := &Future{ch: resChannel}
	work := Work{work: function, args: args, res: future}
	t.taskChannel <- work
	return future
}

// 线程池后台监控程序
func (t *ThreadPool) daemon() {
	if len(t.taskChannel) > t.threshold {
		// 准备在起的协程数
		threadNum := min(t.maxThreads - len(t.workers), len(t.taskChannel) / t.discount)
		for i := 0; i < threadNum; i++ {
			worker := NewWorker(t.taskChannel, false, t.keepAliveTime, t.wg)
			go worker.Work(t.ctx)
			t.workers = append(t.workers, worker)
			t.wg.Add(1)
		}
	}
}

// Close 优雅关闭协程池
func (t *ThreadPool) Close() {
	t.cancel()
	t.wg.Wait()
	close(t.taskChannel)
}