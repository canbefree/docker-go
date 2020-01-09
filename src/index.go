package main

import (
	"log"
	"context"
	"time"
)

func main(){
	ch := make(chan int)
	go func(){
		ch <- 1
		time.Sleep(time.Second * 10)
	}()

	select{
	case <- ch:
		log.Println("get result")
	}
	// run()
}

type k int

func run(){
	const a k = 2
	ctx, cancel := context.WithCancel(context.Background())
	ctx  = context.WithValue(ctx,a,"23")
	go watch(ctx, "12")
	// go watch(ctx, "123")
	time.Sleep(time.Second * 1 )
	cancel()
	time.AfterFunc(5*time.Second,func(){
		log.Println("sfs")
	})
	time.Sleep(time.Second * 7)
}

func watch(ctx context.Context, str string){
	// http请求 10s
	for{
		select{
		case <- ctx.Done():
			log.Println("cancel")
			log.Println(ctx.Value(12))
			return
		}
	}	
}