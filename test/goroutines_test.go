package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMainFuncExit(t *testing.T) {
	Convey("item", t, func() {
		g1()
	})
	time.Sleep(10 * time.Second)
}

func g1() {
	go g2() // 主进程存在就会一直执行
	return
}

func g2() {
	for {
		time.Sleep(time.Second)
		fmt.Println("g2---")
	}
}
