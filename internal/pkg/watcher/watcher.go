package watcher

import (
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

type Watcher struct {
	id   string      // Watcher ID.
	inCh chan string // Input channel.
	//outCh       chan *Counter  // Updates to counter will notify this channel.
	outCh       chan *OutWatcher // Updates to counter will notify this channel.
	counter     *Counter         // The counter.
	counterLock *sync.RWMutex    // Lock for counter.
	quitChannel chan struct{}    // Quit.
	running     sync.WaitGroup   // Run, Amy, Run!
}

func New() *Watcher {
	w := Watcher{}
	w.id = uuid.NewString()
	w.inCh = make(chan string, 1)
	w.outCh = make(chan *OutWatcher, 1)
	w.counter = &Counter{Iteration: 0}
	w.counterLock = &sync.RWMutex{}
	w.quitChannel = make(chan struct{})
	w.running = sync.WaitGroup{}
	return &w
}

// Start watcher in another Go routine, Stop() must be called at the end.
func (w *Watcher) Start() error {
	w.running.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case randStr := <-w.inCh:
				w.counter.Iteration += 1
				select {
				//case w.outCh <- w.counter:
				//outWatcher{randStr: randStr, counter: w.counter}:
				case w.outCh <- &OutWatcher{RandStr: randStr, Counter: w.counter}:
				case <-w.quitChannel:
					return
				}
			case <-w.quitChannel:
				return
			}
		}
	}(&w.running)

	return nil
}

func (w *Watcher) Stop() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	close(w.quitChannel)
	w.running.Wait()
}

func (w *OutWatcher) GetoutWacherStr() ([]byte, error) {
	temp := struct {
		Iteration int    `json:"iteration"`
		Value     string `json:"value"`
	}{
		Iteration: w.Counter.Iteration,
		Value:     w.RandStr,
	}

	return json.Marshal(temp)

}

func (w *Watcher) GetWatcherId() string { return w.id }

func (w *Watcher) Send(str string) { w.inCh <- str }

// func (w *Watcher) Recv() <-chan *Counter { return w.outCh }
func (w *Watcher) Recv() <-chan *OutWatcher { return w.outCh }

func (w *Watcher) ResetCounter() {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()

	w.counter.Iteration = 0

	select {
	//case w.outCh <- w.counter:
	case w.outCh <- &OutWatcher{RandStr: "", Counter: w.counter}:
	case <-w.quitChannel:
		return
	}
}
