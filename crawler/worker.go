package crawler

import (
	"seowl/logger"
	"sync"
)

type workerPool struct {
	size          int
	crawlerEngine *Crawler
	wg            *sync.WaitGroup
}

func InitWorkerPool(crwlr *Crawler) *workerPool {
	return &workerPool{
		size:          poolSize,
		crawlerEngine: crwlr,
		wg:            &sync.WaitGroup{},
	}

}

func (wp *workerPool) work() {

	var log = logger.GetLogger()
	defer wp.wg.Done()

	log.Debug("working")

	// access the queue and parse
	for url := range wp.crawlerEngine.queue {

		log.Debug("processing ", url)
		close(wp.crawlerEngine.queue)
		return
	}
}

func (wp *workerPool) Start() {

	for i := 0; i < wp.size; i++ {
		wp.wg.Add(1)
		go wp.work()
	}

	wp.wg.Wait()

}
