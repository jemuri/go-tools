package pool

import (
	"errors"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const(
	// CLOSED represents that the pool is closed.
	CLOSED = 1
)

type Option func(*Options)

type Options struct {
	ExpiryDuration time.Duration
	PanicHandler func(interface{})
}


// goWorker is the actual executor who runs the tasks,
// it starts a goroutine that accepts tasks and
// performs function calls.
type goWorker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan func()

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}


// run starts a goroutine to repeat the process
// that performs the function calls.
func (w *goWorker) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			if p := recover(); p != nil {
				if w.pool.panicHandler != nil {
					w.pool.panicHandler(p)
				} else {
					log.Printf("worker exits from a panic: %v\n", p)
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					log.Printf("worker exits from panic: %s\n", string(buf[:n]))
				}
			}
		}()

		for f := range w.task {
			if f == nil {
				// 若没有可执行任务,就宣布减少一个worker执行数量
				w.pool.decRunning()
				return
			}
			f()
			if ok := w.pool.revertWorker(w); !ok {
				break
			}
		}
	}()
}

// revertWorker puts a worker back into free pool, recycling the goroutines.
func (p *Pool) revertWorker(worker *goWorker) bool {
	if atomic.LoadInt32(&p.release) == CLOSED || p.Running() > p.Cap() {
		worker.task <- nil
		return false
	}
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)

	// Notify the invoker stuck in 'retrieveWorker()' of there is an available worker in the worker queue.
	p.cond.Signal()
	p.lock.Unlock()
	return true
}



// Pool accept the tasks from client, it limits the total of goroutines to a given number by recycling goroutines.
type Pool struct {
	// capacity of the pool.
	capacity int32

	// running is the number of the currently running goroutines.
	running int32

	// expiryDuration set the expired time (second) of every worker.
	expiryDuration time.Duration

	// workers is a slice that store the available workers.
	workers []*goWorker

	// release is used to notice the pool to closed itself.
	release int32

	// lock for synchronous operation.
	lock sync.Mutex

	// cond for waiting to get a idle worker.
	cond *sync.Cond

	// once makes sure releasing this pool will just be done for one time.
	once sync.Once


	// panicHandler is used to handle panics from each worker goroutine.
	// if nil, panics will be thrown out again from worker goroutines.
	panicHandler func(interface{})
}

func NewPool(size int, option ...Option) *Pool {

	if size == 0 {
		size = 2
	}

	opts := new(Options)
	for _, option := range option {
		option(opts)
	}


	p:= &Pool{
		capacity:         int32(size),
		workers:          make([]*goWorker, 0, size),
		expiryDuration:   opts.ExpiryDuration,
		panicHandler:     opts.PanicHandler,
	}
	p.cond = sync.NewCond(&p.lock)

	return p
}

// Submit submits a task to this pool.
func (p *Pool) submit(task func()) error {

	if w := p.retrieveWorker(); w == nil {
		return errors.New("没有可用的worker")
	}else {
		w.task <- task
	}

	return nil
}
// Running returns the number of the currently running goroutines.
func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

// Free returns the available goroutines to work.
func (p *Pool) Free() int {
	return int(atomic.LoadInt32(&p.capacity) - atomic.LoadInt32(&p.running))
}

// Cap returns the capacity of this pool.
func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}


func (p *Pool) retrieveWorker() *goWorker {
	var w *goWorker
	spawnWorker := func() {
		w = &goWorker{
			pool: p,
			task: make(chan func(), 1),
		}
		w.run()
	}

	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	if n >= 0 {
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
		p.lock.Unlock()
	} else if p.Running()< p.Cap() {
		p.lock.Unlock()
		spawnWorker()
	}


	return w
}

// incRunning increases the number of the currently running goroutines.
func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}


func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

// Release Closes this pool.
func (p *Pool) Release() {
	p.once.Do(func() {
		atomic.StoreInt32(&p.release, 1)
		p.lock.Lock()
		idleWorkers := p.workers
		for i, w := range idleWorkers {
			w.task <- nil
			idleWorkers[i] = nil
		}
		p.workers = nil
		p.lock.Unlock()
	})
}
