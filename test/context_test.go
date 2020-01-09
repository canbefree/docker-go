package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContent(t *testing.T) {
	Convey("test time out sub goroutines can be process: ", t, func() {
		Timeout()
	})
	time.Sleep(3*time.Second) //延长主进程事件  子进程还是执行完毕
}

// 10 秒超时
func Timeout() (Result string, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	return slowOperation(ctx)
}

func slowOperation(ctx context.Context) (Result string, err error){
	// 一种场景是这种  需要协程执行，必须父进程结束或者 在协程加入 ctx控制才能使子进程不执行
	// 还有一种是利用 chan ，pipeline 可以交替执行的 ,使用select来控制，便可以实施结束
	go func(){ //为啥执行不到？？？？ 因为主进程结束了
		time.Sleep(5*time.Second)
		fmt.Println("sub goroutines runs")
	}()
	select{
	case <- ctx.Done():
		err = ctx.Err()
		fmt.Println(ctx.Err())
	}
	return 
}