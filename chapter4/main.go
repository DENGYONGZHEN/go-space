package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {

	// confinement_ad_hoc()
	// confinement_lexical()
	// confinement_lexical2()
	// goroutine_leak()
	// goroutine_leak_mitigate()
	// goroutine_leak_write()
	// goroutine_leak_write_mitigate()
	// or_channel()
	// error_example()
	// error_handle()
	// pipelines()
	// handy_generator()
	// fanin_fanout_naive_example()
	// fanin_fanout_optimize()
	// tee_channel()
	// bridge_channel()
	// useChannelSimulateContext()
	useContext()

}

// 暂时隔离。 按照约定，共享数据不乱用，但是随着人员变动或者时间推移，就会变得很糟糕
func confinement_ad_hoc() {
	data := make([]int, 4) //loopData函数和下面的for循环中都可以访问这个数据，按照约定，只有loopData函数可以访问。但是随着项目的推进，就不一定了

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

// 词法隔离：靠变量作用域强制限制访问
func confinement_lexical() {

	//限制写的操作只存在这个函数的内部的goroutine的作用域内
	chanOwner := func() <-chan int {
		result := make(chan int, 5)

		go func() {
			defer close(result)
			for i := 0; i <= 5; i++ {
				result <- i
			}
		}()
		return result
	}
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

func confinement_lexical2() {

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	//每个 goroutine 只接收到它自己的切片（slice）副本，该切片的作用域只在 goroutine 内部。
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}

func goroutine_leak() {

	dowork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			// defer 也不会执行，completed 也不会被 close，造成 goroutine 泄漏。
			defer close(completed)
			//strings为nil的时候，会永久阻塞，不退出
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	dowork(nil) //这里传入的nil，会伴随着main函数的生命周期，一直存在于内存中，不会执行完成，就会造成内存泄漏
	fmt.Println("Done.")
}

func goroutine_leak_mitigate() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

// goroutine blocked at write operation
func goroutine_leak_write() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int() //在执行3次后，会block
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func goroutine_leak_write_mitigate() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(1 * time.Second)
}

// "or-channel pattern"，主要用来等待多个 channel 中任意一个先结束。
// 多个 goroutine 像一条链串起来，任何一环断开（完成），都会依次向上触发所有 goroutine 退出，确保不泄漏。
func or_channel() {

	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func error_example() {

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhostxx"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

type Result struct {
	Error    error
	Response *http.Response
}

func error_handle() {

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	errCount := 0
	urls := []string{"a", "https://www.google.com", "https://badhostxx", "b", "c"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors,breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}

}

func pipelines() {

	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:

				}
			}
		}()
		return intStream
	}
	// 每一步（stage）接收到的 channel 参数，保证了该 stage 的输入来自上一个 stage，并且输出会被正确传递到下一个 stage。
	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {

		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)

			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream

	}
	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {

		addedStream := make(chan int)
		go func() {
			defer close(addedStream)

			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream

	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

}

func handy_generator() {

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:

					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})

		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)

			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	rand := func() interface{} { return rand.Int() }
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}

func fanin_fanout_naive_example() {
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)

			//检查质数的算法。质数就是大于1的自然数，只能被1和本身整除
			for integer := range intStream {
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}
	rand := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func fanin_fanout_optimize() {
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)

			//检查质数的算法。质数就是大于1的自然数，只能被1和本身整除
			for integer := range intStream {
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:

				}
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream

	}
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")

	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	//根据cpu的数量，创建多个流，把这多个输入流都放入fanIn中，通过fanIn再将这多个输入流中的数据整合到一个流中，在获取10个
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

// 内部 select 的 <-done 后不加 return 是因为：
// 🎯 架构清晰：退出逻辑集中在外层
// 🔄 控制流明确：统一回到主循环处理退出
// 🛡️ 资源安全：确保 defer 语句正确执行
// 📊 状态完整：避免在中间状态退出导致数据不一致
func tee_channel() {

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:

					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})

		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})

		go func() {
			defer close(valStream)

			for {
				select {
				case <-done: // 唯一的退出点
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done: // 回到外层，由外层的 case <-done 处理退出
					}
				}
			}
		}()
		return valStream
	}

	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil //设置为nil后。后续的输入会block
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2

	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}

func bridge_channel() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})

		go func() {
			defer close(valStream)

			for {
				select {
				case <-done: // 唯一的退出点
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done: // 回到外层，由外层的 case <-done 处理退出
					}
				}
			}
		}()
		return valStream
	}

	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {

		valStream := make(chan interface{})

		go func() {
			defer close(valStream)

			for {
				var stream <-chan interface{}

				select {
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}

				//orDone就是为了方便直接可以循环取数据，不用考虑还要监听channel是否关闭。因为都封装在orDone函数中了
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}

func useChannelSimulateContext() {

	locale := func(done <-chan interface{}) (string, error) {
		select {
		case <-done:
			return "", fmt.Errorf("canceled")
		case <-time.After(1 * time.Minute):
		}
		return "EN/US", nil
	}

	genGreeting := func(done <-chan interface{}) (string, error) {
		switch locale, err := locale(done); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "hello", nil
		}
		return "", fmt.Errorf("unsupported locale")
	}
	genFarewell := func(done <-chan interface{}) (string, error) {
		switch locale, err := locale(done); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "goodbye", nil
		}
		return "", fmt.Errorf("unsupported locale")
	}

	printGreeting := func(done <-chan interface{}) error {
		greeting, err := genGreeting(done)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", greeting)
		return nil
	}
	printFarewell := func(done <-chan interface{}) error {
		farewell, err := genFarewell(done)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", farewell)
		return nil
	}

	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Wait()
}

func useContext() {

	locale := func(ctx context.Context) (string, error) {
		if deadline, ok := ctx.Deadline(); ok {
			if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
				return "", context.DeadlineExceeded
			}
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(1 * time.Minute):
		}
		return "EN/US", nil
	}

	genGreeting := func(ctx context.Context) (string, error) {
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		switch locale, err := locale(ctx); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "hello", nil
		}
		return "", fmt.Errorf("unsupported locale")
	}
	genFarewell := func(ctx context.Context) (string, error) {
		switch locale, err := locale(ctx); {
		case err != nil:
			return "", err
		case locale == "EN/US":
			return "goodbye", nil
		}
		return "", fmt.Errorf("unsupported locale")
	}

	printGreeting := func(ctx context.Context) error {
		greeting, err := genGreeting(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", greeting)
		return nil
	}
	printFarewell := func(ctx context.Context) error {
		farewell, err := genFarewell(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%s world!\n", farewell)
		return nil
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			return
		}
	}()

	wg.Wait()
}
