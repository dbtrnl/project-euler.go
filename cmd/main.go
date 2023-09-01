package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/dbtrnl/project-euler/golang/internal/problems"
)

var problemsSlice = []func() int{
	problems.Problem1,
	problems.Problem2,
	problems.Problem3,
	problems.Problem4,
	problems.Problem5,
	problems.Problem6,
	problems.Problem7,
	problems.Problem8,
	problems.Problem9,
	problems.Problem10,
	problems.Problem11,
	problems.Problem12,
	problems.Problem13,
	problems.Problem14,
	problems.Problem15,
	problems.Problem16,
	problems.Problem17,
	problems.Problem18,
}

func main() {
	for _, fn := range problemsSlice {
		result := fn()
		fnName := strings.Split(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")[2]
		fmt.Printf("Answer of '%s' is: %d\n", fnName, result)
	}
}
