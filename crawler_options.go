package crawl

import "net/http"

// Option - Crawl option.
type Option func(*crawl)

// options - Crawl options.
type options struct {
	concurrency   int
	queueCapacity int
	headers       map[string]string
}

// WithTransport - Sets crawl HTTP transport.
func WithTransport(transport *http.Transport) Option {
	return func(crawl *crawl) {
		crawl.transport = transport
	}
}

// WithQueue - Sets crawl queue.
// Default: creates queue using NewQueue() with capacity of WitWithQueueCapacity().
func WithQueue(queue Queue) Option {
	return func(crawl *crawl) {
		crawl.queue = queue
	}
}

// WithDefaultHeaders - Sets crawl default headers.
// Default: empty.
func WithDefaultHeaders(headers map[string]string) Option {
	return func(crawl *crawl) {
		crawl.opts.headers = headers
	}
}

// WithUserAgent - Sets crawl default user-agent.
func WithUserAgent(ua string) Option {
	return func(crawl *crawl) {
		if crawl.opts.headers == nil {
			crawl.opts.headers = make(map[string]string)
		}
		crawl.opts.headers["User-Agent"] = ua
	}
}

// WithConcurrency - Sets crawl concurrency.
// Default: 1000.
func WithConcurrency(n int) Option {
	return func(crawl *crawl) {
		crawl.opts.concurrency = n
	}
}

// WithQueueCapacity - Sets queue capacity.
// It sets queue capacity if a queue needs to be created and it sets a capacity of channel in-memory queue.
// It also sets capacity of errors buffered channel.
// Default: 10000.
func WithQueueCapacity(n int) Option {
	return func(crawl *crawl) {
		crawl.opts.queueCapacity = n
	}
}
