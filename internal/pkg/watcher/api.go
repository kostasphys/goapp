package watcher

type Counter struct {
	Iteration int `json:"iteration"`
}

type OutWatcher struct {
	RandStr string
	Counter *Counter
}

type CounterReset struct{}
