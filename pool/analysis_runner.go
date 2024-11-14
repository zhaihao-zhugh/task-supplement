package pool

import (
	"gpk/logger"
	"supplementary-inspection/model"
	"sync"
	"time"
)

var PAnalysisRunner *AnalysisRunner

func Run() {
	PAnalysisRunner = &AnalysisRunner{
		Items: make([]model.AnalysisItem, 0),
		ch:    make(chan struct{}, 1),
	}
	PAnalysisRunner.Run()
}

// AnalysisRunner 分析执行者
type AnalysisRunner struct {
	sync.Mutex
	// wg      sync.WaitGroup
	ch      chan struct{}        // channel to limited number of goroutines
	Items   []model.AnalysisItem // 待分析元素数组
	Workers sync.Map
}

func NewAnalysisRunner() *AnalysisRunner {
	return &AnalysisRunner{
		Items: make([]model.AnalysisItem, 0),
		ch:    make(chan struct{}, 1),
	}
}

func (runner *AnalysisRunner) Append(item model.AnalysisItem) {
	runner.Lock()
	defer runner.Unlock()

	runner.Items = append(runner.Items, item)
}

func (runner *AnalysisRunner) Run() {
	for {
		repeat := make(map[string]uint)
		length := 0

		for _, item := range runner.Items {
			// 限制最大请求数量
			// if i >= Settings.Default.RequestLength {
			// 	break
			// }

			_, ok := repeat[item.ObjectID]

			// 待发送请求总长度
			if ok {
				break
			} else {
				repeat[item.ObjectID] = 0
			}

			length += 1
		}

		if length > 0 {
			items := runner.Items[0:length]
			worker := NewAnalysisWorker()
			runner.Workers.Store(worker.RequestID, worker)
			runner.ch <- struct{}{}
			logger.Infof("===>%s", worker.RequestID)
			go func() {
				_ = worker.Work(runner.ch, items)
				runner.Workers.Delete(worker.RequestID)
			}()

			runner.Lock()
			runner.Items = runner.Items[length:]
			runner.Unlock()
		}

		time.Sleep(time.Second)
	}
}
