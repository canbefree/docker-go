package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	. "github.com/smartystreets/goconvey/convey"
)

// chan 应该在协程里面关闭还是外部直接关闭？ 创建了chan 然后关闭？？ 

func TestSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			ch <- rand.Int()
		}
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	}
	select {
	case v := <-ch:
		fmt.Println(v)
	}

}

func TestClose(t *testing.T) {
	Convey("err",t,func(){
		defer func(){
			err := recover()
			if err != nil {
				So(err, ShouldEqual, "send on closed channel")
			}
		}()
		ch := make(chan int)
		go func(){
			close(ch)
		}()
		ch <- 35		
	})

	// ch <- 34
}

func intoChan() {
	for {
		time.Sleep(time.Second)
		// a := rand.Int()
	}
}
