package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1.6.8 context包
// https://www.cnblogs.com/lchb/articles/10844423.html
// 在Go1.7发布时，标准库增加了一个context包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作，
// 官方有博文对此做了专门介绍。我们可以用context包来重新实现前面的线程安全退出或超时的控制:

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}

// 当并发体超时或main主动停止工作者Goroutine时，每个工作者都可以安全退出。

// Go语言是带内存自动回收特性的，因此内存一般不会泄漏。在前面素数筛的例子中，GenerateNatural和PrimeFilter函数内部都启动了新的Goroutine，当main函数不再使用管道时后台Goroutine有泄漏的风险。我们可以通过context包来避免这个问题，下面是改进的素数筛实现：

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	// 通过 Context 控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())

	ch := GenerateNatural(ctx) // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime) // 基于新素数构造的过滤器
	}

	cancel()
}

// 当main函数完成工作前，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。

// 并发是一个非常大的主题，我们这里只是展示几个非常基础的并发编程的例子。官方文档也有很多关于并发编程的讨论，
// 国内也有专门讨论Go语言并发编程的书籍。读者可以根据自己的需求查阅相关的文献。

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	eatNum := chiHanBao(ctx)

	for n := range eatNum {
		fmt.Println("-->> Eat:", n)
		if n >= 10 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果。。。")
	time.Sleep(1 * time.Second)
}

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	// n := 0
	// 时间
	t := 0
	go func() {
		n := 1
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("===>>> 耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
				time.Sleep(1*time.Second)
			}
		}
	}()
	
	fmt.Println("------------ begin ----------")
	return c
}
