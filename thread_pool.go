package gulc


import (
	"sync"
	"time"
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
	quit        chan struct{}
	keepAliveTime  time.Duration
	wg *sync.WaitGroup
}

func NewWorker(taskChannle chan Work, isCore bool, quit chan struct{}, keepAliveTime time.Duration, wg *sync.WaitGroup) *Worker {
	return &Worker{taskChannel: taskChannle, isCore: isCore, quit: quit, keepAliveTime: keepAliveTime, wg: wg}
}


func (w *Worker) Work() {
	defer w.wg.Done()
	timer := time.NewTimer(w.keepAliveTime)
	for {
		select {
		case <- timer.C:
			if !w.isCore {
				return
			}
		case <- w.quit: 
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
	quit chan struct{}
	threshold int
	discount int
	wg *sync.WaitGroup
}

func NewThreadPool(coreThreads, maxThreads int, keepAliveTime time.Duration, threshold, discount int) *ThreadPool {
	taskChannel := make(chan Work, maxThreads * 2 + 1)
	quit := make(chan struct{}, maxThreads)
	workers := make([]*Worker, 0)
	var wg sync.WaitGroup
	for i := 0; i < coreThreads; i++ {
		worker := NewWorker(taskChannel, true, quit, keepAliveTime, &wg)
		go worker.Work()
		workers = append(workers, worker)
		wg.Add(1)
	}
	t := &ThreadPool{workers: workers, coreThreads: coreThreads, maxThreads: maxThreads, keepAliveTime: keepAliveTime, taskChannel: taskChannel, quit: quit, threshold: threshold, discount: discount, wg: &wg}
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
			worker := NewWorker(t.taskChannel, false, t.quit, t.keepAliveTime, t.wg)
			go worker.Work()
			t.workers = append(t.workers, worker)
			t.wg.Add(1)
		}
	}
}

// Close 优雅关闭协程池
func (t *ThreadPool) Close() {
	for i := 0; i < t.maxThreads; i++ {
		t.quit <- struct{}{}
	}
	t.wg.Wait()
	close(t.quit)
	close(t.taskChannel)
}