package main

import (
	"fmt"
	"sync"
	"time"
)

/*
select + channel实现非阻塞的互斥锁
*/
type MyMutex struct{
	c chan struct{}
}

func NewMyMutex() MyMutex {
	cc := MyMutex{
		c: make(chan struct{}, 1),
	}
	return cc
}

func (c *MyMutex) Lock(tt time.Duration, gid int) bool {
	t := time.NewTimer(tt)
	select {
		case <-t.C :
			fmt.Println("g", gid, " lock failed")
			return false
		case c.c <- struct{}{} :
			fmt.Println("g", gid, " lock success")
			return true
	}
}

func (c *MyMutex) UnLock(tt time.Duration, gid int) bool {
	t := time.NewTimer(tt)
	select {
	case <-t.C:
		fmt.Println("g", gid, " unlock failed")
		return false
	case <-c.c:
		fmt.Println("g", gid, " unlock success")
		return true
	}
}

func main() {
	mylock := NewMyMutex()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		println("g1 getting lock..")
		islock := mylock.Lock(time.Second * 3, 1)
		defer mylock.UnLock(time.Second , 1)
		if islock {
			println("process 1 going...")
			println("process 1 done")
			time.Sleep(time.Second)
		}
		wg.Done()
	} ()

	go func() {
		println("g2 getting lock..")
		islock := mylock.Lock(time.Second * 3, 2)
		defer mylock.UnLock(time.Second, 2)
		if islock {
			println("process 2 going...")
			println("process 2 done")
			time.Sleep(time.Second * 5)
		}
		wg.Done()
	} ()

	go func() {
		println("g3 getting lock..")
		islock := mylock.Lock(time.Second * 3, 2)
		defer mylock.UnLock(time.Second, 2)
		if islock {
			println("process 3 going...")
			println("process 3 done")
			time.Sleep(time.Second)
		}
		wg.Done()
	} ()

	wg.Wait()
}