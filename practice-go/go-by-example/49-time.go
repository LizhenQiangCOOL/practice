package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	//当前时间
	now := time.Now()
	p(now)

	//构造时间
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651389237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month()) //返回得是英文月份，非数字
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Nanosecond())
	p(then.Location()) //时间区域

	p(then.Weekday())

	p(then.Before(now)) //两个时间的比较
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) //时间之差
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

}
