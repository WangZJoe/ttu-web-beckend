package logRecorder

import (
	"fmt"
	"log"
	"runtime"
)

func PrintmynameTest() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func TestFunc() {
	println("test func")
}

func Log_test() {
	print("\n\n\nstart log test \n")
	fmt.Println("Hello, world!")
	log.Println("log test")

	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	var s1 string
	s1 = "output sample"
	log.Println(s1)
	//time.Sleep(3 * time.Second)
	log.Println("log test")
	CheckFile(s1)
	CreateFile()
}

func Foo() {
	fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
