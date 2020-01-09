package test

import (
	"log"
	"testing"
)

func TestPanicAndRecover(t *testing.T){
	defer func(){
		err := recover()
		if err != nil{
			log.Println(err)
			t.Log(err)
		}
	}()
	defer func(){
		t.Log("defer是倒序执行的")
	}()

	panic("error")
	t.Error("其后面的代码不会执行")
}
