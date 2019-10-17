package main

import (
	"fmt"
	"function"
	"time"
)

func send(a chan int) {
	for i := 0; i < 1; i++ {
		a <- i
	}
}

func recv(b chan int, index int) {
	for {
		select {
		case x := <-b:
			fmt.Println("index=", index, " x=", x)
		}
	}
}

type MsgCreator struct {
	msgCh  chan string
	name   string
	closed int
}

type MsgRecver struct {
	msgCh  chan string
	name   string
	userId int32
}

type Dispatcher struct {
	c *MsgCreator
	s [](*MsgRecver)
}

func (c *MsgCreator) sendMsg(msg string) {
	if c.closed == 0 {
		c.msgCh <- msg
	}
}

func (r *MsgRecver) recvMsg() {
	for {
		msg, ok := <-r.msgCh
		if ok {
			fmt.Println("recver:", r.name, " user_id:", r.userId, " msg:", msg)
		} else {
			fmt.Println("recver:", r.name, " user_id:", r.userId, " channel closed.")
			break
		}
	}
}

func NewDispatcher(c *MsgCreator) *Dispatcher {
	d := new(Dispatcher)
	d.c = c
	d.s = make([]*MsgRecver, 0, 10)

	return d
}

func (d *Dispatcher) AddRecver(r *MsgRecver) {
	d.s = append(d.s, r)
}

func (d *Dispatcher) Observer() {
	fmt.Println("observer:", d.c.name)

	for {
		msg, ok := <-d.c.msgCh
		if ok {
			fmt.Println("dispatch Msg:", msg)
			for k, v := range d.s {
				fmt.Println("dispatch:", k, " name:", v.name)
				v.msgCh <- msg
			}
		} else {
			fmt.Println("observer:", d.c.name, " exit.")
			for _, v := range d.s {
				close(v.msgCh)
			}
			break
		}
	}
}

func main() {
	ch := make(chan int)

	go recv(ch, 1)
	go recv(ch, 2)
	// go recv(ch, 3)
	// go recv(ch, 4)

	go send(ch)
	go send(ch)
	go send(ch)

	time.Sleep(1 * time.Second)
	// select {}

	fmt.Println("-------- observer test ----------")

	// sender and recver new
	c := &MsgCreator{make(chan string), "tom", 0}
	r1 := &MsgRecver{make(chan string), "user1", 1}
	r2 := &MsgRecver{make(chan string), "user2", 2}
	r3 := &MsgRecver{make(chan string), "user3", 3}
	r4 := &MsgRecver{make(chan string), "user4", 4}

	// recv goroutine
	go r1.recvMsg()
	go r2.recvMsg()
	go r3.recvMsg()
	go r4.recvMsg()

	obs := NewDispatcher(c)

	// add recv to observer
	obs.AddRecver(r1)
	obs.AddRecver(r2)
	obs.AddRecver(r3)
	obs.AddRecver(r4)

	// send msg
	go func() {
		for i := 0; i < 1000; i++ {
			msg := fmt.Sprintf("msg %d", i)
			c.sendMsg(msg)
			time.Sleep(10 * time.Millisecond)
		}

		close(c.msgCh)
	}()

	// observer
	obs.Observer()

	time.Sleep(1 * time.Second)

	//fmt.Println("mysql ip:", AA)
	function.Test()
}
