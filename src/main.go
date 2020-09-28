package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "study-demo/src/classtwo"
	"time"
)

/**
 * @Author: gaoz
 * @Date: 2020/9/28
 */

func main() {

	// 强类型
	var num int
	num = 1
	fmt.Println(num)

	name := "li"
	fmt.Println(name)

	// 下面启动三个goroutine
	// 执行顺序是乱的
	go func() {
		fmt.Println("in goroutine-1")
	}()
	go func() {
		fmt.Println("in goroutine-2")
	}()
	go func() {
		fmt.Println("in goroutine-3")
	}()

	// jeep
	jeep := Jeep{}
	jeep.Running()

	RecordInfo(FileLogger{}, "log....info")

	// sleep 3秒 防止上面的协程执行没有完成
	time.Sleep(3 * time.Second)

	fmt.Print("hello go")


	//classtwo.HttpStart()
}


//------------------------
// 结构体
// 想象成对象即可
type Car struct {
	// 成员
	Tire string
}

// 成员方法
func (car Car) Running() {
	fmt.Println("is running...")
}

type Jeep struct {
	// 继承
	Car
}

//------------------------

// 定义一个接口
type Logger interface {
	Info(msg string)
	Error(msg string)
}

type FileLogger struct {
}

// fileLogger实现了Info、Error
// 就实现了 Logger
func (fileLogger FileLogger) Info(msg string) {
	fmt.Println("msg write file..."+ msg)
}
func (fileLogger FileLogger) Error(msg string) {
	fmt.Println("msg error file... "+ msg)
}

//------------------------


/**

1、参数及返回值需指定类型
2、大括号不能换行
*/
func getName(id int) string {
	fmt.Println(id)
	return "liu"
}

func RecordInfo(logger Logger, msg string) {
	logger.Info(msg)
}