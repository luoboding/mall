package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type Site struct {
	url  string
	size int
}

func main() {
	channelTest()
	// condTest()
	getSysInfo()
}

func getSysInfo() {
	fmt.Println(runtime.NumCPU())
}

// 获取url内容的长度
func channelTest() {
	cha := []chan Site{
		make(chan Site),
		make(chan Site),
	}
	go pageSize("http://www.baidu.com", cha[0])
	go pageSize("http://www.sogou.com", cha[1])
	// a := <-cha
	// b := <-cha
	for {
		select {
		case e := <-cha[0]:
			fmt.Println("url is ", e.url, " content size is ", e.size)
		case a := <-cha[1]:
			fmt.Println("url is ", a.url, " content size is ", a.size)
	}

	// fmt.Println("url is ", a.url, " content size is ", a.size)
	// fmt.Println("url is ", b.url, " content size is ", b.size)
}

func pageSize(url string, c chan Site) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	c <- Site{
		url:  url,
		size: len(content),
	}
}

// cond test start
var productCount int

type Factory struct {
	cond  *sync.Cond
	goods []int
}

func (self *Factory) init() {
	self.cond = sync.NewCond(&sync.Mutex{})
}

func (self *Factory) product() {
	self.cond.L.Lock()
	defer self.cond.L.Unlock()
	productCount++
	self.goods = append(self.goods, productCount)
	self.cond.Signal()
	fmt.Println("生产数据完成，发信号唤醒最早Wait的")
}

func (self *Factory) consume() {
	self.cond.L.Lock()
	defer self.cond.L.Unlock()
	for {
		if len(self.goods) == 0 {
			fmt.Println("没有可以消费的，sleep")
			self.cond.Wait()
			fmt.Println("wake: 开始消费")
		} else {
			break
		}
	}
	self.goods = self.goods[1:]
}

func (self *Factory) line1() {
	for {
		self.consume()
		time.Sleep(time.Second * 1)
	}
}

func (self *Factory) line2() {
	for {
		self.product()
		time.Sleep(time.Second * 1)
	}
}

func condTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fa := &Factory{}
	fa.init()
	go fa.line2()
	go fa.line1()
	select {}
}
