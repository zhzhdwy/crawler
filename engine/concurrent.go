package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// 并发点主要在于Fetcher中，这个方法已经被simple.go中包装成了worker
func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 定义两个chan，in用于放任务，worker们从in中抢任务
	// out用于worker中结果的传入，当然也是抢。
	// 因为我只生成了一进一出的两个chan
	//in := make(chan Request)
	out := make(chan ParseResult)

	// new
	e.Scheduler.Run()

	// 这个地方主要是将in赋值给Scheduler中的workerchan
	// 其实workerchan就是in
	//e.Scheduler.ConfigureMasterWorkChan(in)

	// 生成若干个worker
	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)

		//new
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}

	}

}

var urls = make(map[string]interface{})

func isDuplicate(url string) bool {

	if _, ok := urls[url]; ok {
		//log.Printf("%v \n %v", urls, url)
		return true
	} else {
		urls[url] = 1
		return false
	}
}

// 为什么要有createworker呢，主要是worker是实际工作的方法
// 而createworker中主要是生成，并且规划chan的使用
//func createWorker(in chan Request, out chan ParseResult) {
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// new
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
