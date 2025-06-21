package crawler

type Crawler struct {
	inputURL string
	visited  map[string]bool
	queue    chan string
}

func NewCrawler(url string) *Crawler {

	c := Crawler{
		inputURL: url,
	}

	c.queue = make(chan string, 10)
	c.visited = make(map[string]bool)

	return &c
}
