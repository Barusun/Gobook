//对比三个函数的执行效率
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
//使用传统的for循环遍历的方式
func echo1() string {
	var s, sep string
	for i:=1; i<len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}
//使用range 切片的方式
func echo2() string {
	s, sep := "", ""
	for _,arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	return s
}
//使用strings的Join函数
func echo3() string{
	return fmt.Sprintln(strings.Join(os.Args[1:]," "))
}
func main() {
	//使用数组调用循环调用三个echo函数
	var echo = [3]func() string{echo1,echo2,echo3}
	for _, echo_func := range echo{
		start := time.Now()
		s := echo_func()
		ecs := time.Since(start).Nanoseconds()
		fmt.Printf("%d, %s\n",ecs,s)
	}
}


