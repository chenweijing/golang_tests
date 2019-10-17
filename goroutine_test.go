package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
// 一段耗时的计算函数
func consumer(ch chan int) {
	// 无限获取数据的循环
	for {
		// 从通道获取数据
		data := <-ch
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}
	fmt.Println("goroutine exit")
}
func TestGoroutine(t *testing.T) {
	// 传递数据用的通道
	ch := make(chan int)
	for {
		// 空变量, 什么也不做
		var dummy string
		// 获取输入, 模拟进程持续运行
		fmt.Scan(&dummy)
		if dummy == "quit" {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- i
			}
			continue
		}
		// 启动并发执行consumer()函数
		go consumer(ch)
		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
*/

func TestJson(t *testing.T) {
	b := []byte("{\"Name\": \"Wednesday\", \"Age\": 6, \"Parents\": [\"Gomez\", \"Morticia\"]}")
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		fmt.Printf("err %s", err.Error())
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "# is string", vv)

		case int:
			fmt.Println(k, "## is int", vv)

		case []interface{}:
			fmt.Println(k, "### is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}

		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}

func Test1(t *testing.T) {
	// 创建通道
	ch := make(chan int)

	// 关闭通道
	close(ch)

	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))

	// ch <- 1
}

func TestChanne2(t *testing.T) {
	// 创建通道
	ch := make(chan int, 2)
	ch <- 0
	ch <- 1

	// 关闭通道
	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}

	// ch <- 1
}
