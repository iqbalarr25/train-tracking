package helper

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func getCallerInfo(skip int) (string, string, int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return "", "", 0
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // The Base function returns the last element of the path
	return funcName, fileName, lineNo
}

func Exception(err error, other ...interface{}) {
	function, file, line := getCallerInfo(2)
	fmt.Println("ERROR ", time.Now())
	fmt.Println("Function: ", function)
	fmt.Println("File:     ", file)
	fmt.Println("Line:     ", line)
	fmt.Println("Error:    ", err.Error())
	if len(other) > 0 {
		fmt.Println("Other:    ", other[0])
	}
}
