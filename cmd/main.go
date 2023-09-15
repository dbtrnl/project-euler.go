package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/dbtrnl/project-euler.go/internal/problems"
)

var wg sync.WaitGroup
var problemsSlice = []func(*sync.WaitGroup, chan<- int){
	problems.Problem1, problems.Problem2, problems.Problem3, problems.Problem4, problems.Problem5, problems.Problem6,
	problems.Problem7, problems.Problem8, problems.Problem9, problems.Problem10, problems.Problem11, problems.Problem12,
	problems.Problem13, problems.Problem14, problems.Problem15, problems.Problem16, problems.Problem17, problems.Problem18,
	problems.Problem19, problems.Problem20, problems.Problem21, problems.Problem22, problems.Problem23, // problems.Problem24,
	problems.Problem25, // problems.Problem26, // problems.Problem27, // problems.Problem28, // problems.Problem29, // problems.Problem30,
	problems.Problem67,
}
var probLen = len(problemsSlice)
var answerChan = make(chan int, probLen)

func addWaitgroup() {
	wg.Add(probLen)
}

func asynchronousMain() {
	for _, fn := range problemsSlice {
		go func(fn func(*sync.WaitGroup, chan<- int)) {
			fnName := strings.Split(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")[3]
			fn(&wg, answerChan)
			fmt.Printf("%s answer is: %d\n", fnName, <-answerChan)
		}(fn)
	}
	wg.Wait()
}

func synchronousMain() {
	for _, fn := range problemsSlice {
		fn(&wg, answerChan)
		fnName := strings.Split(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")[3]
		fmt.Printf("Answer of '%s' is: %d\n", fnName, <-answerChan)
	}
}

func main() {
	addWaitgroup()
	startTime := time.Now()
	// synchronousMain()
	asynchronousMain()
	elapsedTime := time.Now().Sub(startTime).Seconds()
	fmt.Printf("-----\n%v functions executed in %.2f seconds\n", probLen, elapsedTime)

}
