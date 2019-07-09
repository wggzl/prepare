package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)

	//每一分钟执行一次
	//if expr, err = cronexpr.Parse("* * * * *"); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//每5分钟执行一次
	//if expr, err = cronexpr.Parse("*/6 * * * *"); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	if expr, err = cronexpr.Parse("* * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	//当前时间
	now = time.Now()
	//下一次调度售价仅
	nextTime = expr.Next(now)


	//expr = expr
	//fmt.Println(now, nextTime)

	//等待定时器超时
	//time.NewTimer(nextTime.Sub(now))

	time.AfterFunc(nextTime.Sub(now), func(){
		fmt.Println("被调度了：", nextTime)
	})

	time.Sleep(5 * time.Second)

}
