package main

import (
	"errors"
	"fmt"
)

// 正常情况下，error为nil
// 方法一： 返回个errors.New()
func f1(arg int) (int, error){
	if arg == 42{
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// 方法二： 定义个Error结构体，并实现Error属性（）
// 返回一个 Error结构体指针
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42{
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main()  {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e!=nil{
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, j := range []int{7, 42}{
		if r, e := f2(j); e!=nil {
			fmt.Println("f2 failed:", e)
		}else{
			fmt.Println("f2 worked:", r)
		}
	}

	_, e:=f2(42)
	fmt.Println("e value:", e)
	//e.(*argError) 暂时不知道该操作
	if ae, ok := e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
