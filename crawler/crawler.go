package crawler

const poolSize int = 10

type Crawler struct {
	inputURL string
	visited  map[string]bool
	queue    chan string
}

func NewCrawler(url string) *Crawler {

	c := Crawler{
		inputURL: url,
	}

	c.queue = make(chan string, poolSize)
	c.visited = make(map[string]bool)

	return &c
}

func (c *Crawler) Start() {

	// create worker pool
	wp := InitWorkerPool(c)

	c.visited[c.inputURL] = true
	c.queue <- c.inputURL
	wp.Start()

	// for url := range c.queue {
	// 	// a worker takes the queue and url, parses and puts inside queue
	// }

}
