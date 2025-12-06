package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {

	//go 会创建一个协程
	// go sayHello()

	//匿名函数
	// go func() {
	// 	fmt.Println("hello")
	// }()

	// sayHello2 := func() {
	// 	fmt.Println("hello")
	// }
	// go sayHello2()

	// closure()
	// closureLoop()
	// closureLoop2()
	// goroutine_GC()
	// waitGroup()
	// mutex()
	// rwMutex()
	// cond()
	// cond_broadcast()
	// sync_Once()
	// sync_Once_deadlock()
	// sync_Pool()
	// sync_Pool2()
	// channel2()
	// channel_close()
	// channel_range()
	// channel_unblock_multi_goroutine()
	// channel_buffered()
	// channel_default_value()
	// channel_owner()
	// sync_select()
	// select_random()
	// select_never()
	// select_default()
	select_default_usage()
}

func sayHello() {
	fmt.Println("hello")
}

func closure() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func closureLoop() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

// 一样的结果
func closureLoop2() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func goroutine_GC() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func waitGroup() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	// first step
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	// second step,wait count to be 0
	wg.Wait()

	// third step
	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
func printSomething(s string, wg *sync.WaitGroup) {
	// third step
	defer wg.Done()
	fmt.Println(s)
}

func mutex() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

// 多个持有读锁的协程可以同时读。写锁是排斥的
func rwMutex() {

	//模拟写入的协程
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1 * time.Microsecond)
		}
	}

	//模拟读取的协程，l可以为普通锁或者读写锁
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()

		//开启一个写者协程和多个读者协程
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		//第一个测得是使用读锁的时间，第二个测的是普通锁的时间
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}
}

func cond() {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)

	removerFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		c.L.Unlock()
		c.Signal() //找到等待时间最长的那个协程，并唤醒
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait() //挂起协程，等待唤醒
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removerFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

type Button struct {
	Clicked *sync.Cond
}

func cond_broadcast() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			//这里作为信号，表示这个goroutine已经启动了，所以这里不需要使用defer
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		//主goroutine在这里等待上面的协程启动后继续执行
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}

// sync.Once 记录的是 是否已经执行过一次，而不是记录具体的函数。
// 第一次调用 Do 时传入的函数会执行，之后所有 Do 调用都会被忽略。
// 如果你想不同函数都只执行一次，需要为每个函数分别创建一个 sync.Once 实例。
func sync_Once() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment) //once.Do(f) 保证函数 f 只执行一次，即使多个 goroutine 同时调用 Do
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func sync_Once_deadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // this call can't proceed until "onceA.Do(initA)" returns
	onceA.Do(initA)                    //这个操作或获取锁，在结束前释放锁。
}

func sync_Pool() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func sync_Pool2() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		//当Get获取不到时，调用New创建一个新的
		New: func() interface{} {

			//记录实际创建了几个
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			//Assume something interesting, but quick is being done with this memory.
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}

// func channel() {

// 	var receiveChan <-chan interface{}
// 	var sendChan chan<- interface{}

// 	dataStream := make(chan interface{})
// 	//无方向channel可以直接赋值给有方向channel
// 	receiveChan = dataStream
// 	sendChan = dataStream
// }

func channel2() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutaion, ok := <-stringStream //等待直到channel中存在数据
	fmt.Printf("(%v): %v\n", ok, salutaion)
}
func channel_close() {
	stringStream := make(chan int)
	close(stringStream)
	salutaion, ok := <-stringStream //等待直到channel中存在数据
	fmt.Printf("(%v): %v\n", ok, salutaion)
}

func channel_range() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// 可以使用 range 从 channel 连续接收数据。
	// 当 channel 被关闭并且所有数据都被接收完后，for-range 循环会自动结束。
	// 注意：如果 channel 未被关闭，range 循环会一直阻塞等待新数据。
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

func channel_unblock_multi_goroutine() {

	//无缓存channel，就是长度为0
	begin := make(chan interface{})

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func channel_buffered() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Produce Done.")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func channel_default_value() {
	var dataStream chan interface{}
	//尝试从一个nil channel中获取数据，会引起死锁
	// <-dataStream

	//当尝试向一个nil channel中写入数据，会引起死锁
	// dataStream <- struct{}{}

	//当尝试关闭这个nil channel时，会panic，另外重复关闭一个已经关闭的channel，也会panic
	close(dataStream)
}

// 明确职责
func channel_owner() {

	//channel 的owner应该负责创建和写入和关闭
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")

}

func sync_select() {

	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func select_random() {

	// 这里创建了两个 已关闭的 channel。
	// 已关闭的 channel 总是可读，读取不会阻塞，会立即返回零值。
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	// 如果 select 中有多个 case 同时就绪（可以立即执行），Go 会随机选择一个 执行。
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

}

func select_never() {
	var c <-chan int
	select {
	case <-c: // nil channel 会一直blocked
	case <-time.After(1 * time.Second):
		fmt.Println("Time out.")
	}
}

func select_default() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

func select_default_usage() {

	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

func select_empty() {

	//this statement will simply block forever
	select {}
}
