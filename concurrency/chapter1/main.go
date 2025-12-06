package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// raceCondition()
	// memoryAccessSyn()
	// deadlock()
	// livelock()
	starvation()
}

func raceCondition() {
	var data int
	go func() {
		data++
	}()

	//there is no guarantee what the data is
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

func memoryAccessSyn() {
	var memoryAccess sync.Mutex
	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	//it still not sure the order of the operation,and have performance problems
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}

type value struct {
	mu    sync.Mutex
	value int
}

func deadlock() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)

	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

// 必须在持有锁时调用 Wait()。
// 否则会 panic 或导致竞态。Wait() 期望调用者已经持有 Cond.L。
// Wait() 会释放锁并阻塞，唤醒时会重新获取锁。
// 这一点非常关键：它保证在被唤醒之后，线程能在持锁的情况下检查或修改共享状态。
// Wait() 可能出现“虚假唤醒（spurious wakeup）”，所以总是在循环中检查条件：

// c.L.Lock()
// for !predicate() {
//     c.Wait()
// }
// // 此时 predicate() == true
// c.L.Unlock()

// 也就是说不要只写 if，要写 for。

// 唤醒 API：
// Signal()：唤醒等待队列中的一个 goroutine（任意一个，通常 FIFO/implementation dependent）。
// Broadcast()：唤醒队列中所有 goroutine。
func livelock() {
	//sync.Cond 是 Go 的条件变量（Condition Variable）。这里的作用是让多个 goroutine 以同样的节奏（cadence）“同时行动”。
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		// 每隔 1 毫秒
		for range time.Tick(1 * time.Millisecond) {
			// Broadcast()：唤醒所有正在 Wait() 的 goroutine。
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock() //获取锁
		// Wait()：在条件满足前阻塞当前 goroutine（会释放锁并进入等待，等被唤醒后重新获取锁再返回）。
		cadence.Wait()
		cadence.L.Unlock() //释放锁
	}

	//dir 表示这条路上有几个人（左路或右路）。
	// 当某人尝试走这条路时：
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1) // ① 表示“我想往这个方向走”，先占一个名额
		takeStep()              // ② 等待节奏信号（相当于同步点）

		if atomic.LoadInt32(dir) == 1 { // ③ 检查是不是“只有我一个人走这边”
			fmt.Fprint(out, ". Success!") // ✅ 如果是，就成功
			return true
		}

		takeStep()               // ④ 再次等待同步点
		atomic.AddInt32(dir, -1) // ⑤ 如果不是（说明两个人都选了同一个方向），那我退一步
		return false             // ❌ 表示这次尝试失败
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}

// Polite worker was able to execute 642241 work loop.
// Greedy worker was able to execute 1410789 work loops
func starvation() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begain := time.Now(); time.Since(begain) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int

		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loop.\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}
