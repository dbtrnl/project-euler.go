package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	problems "github.com/dbtrnl/project-euler/golang/internal/problems/1-10"
)

var problemsSlice = []func()int {
	problems.Problem1,
	problems.Problem2,
	problems.Problem3,
	problems.Problem4,
	problems.Problem5,
}

func main() {
	for _, fn := range problemsSlice {
		result := fn()
		fnName := strings.Split(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")[2]
		fmt.Printf("Answer of '%s' is: %d\n", fnName, result)
	}
}

