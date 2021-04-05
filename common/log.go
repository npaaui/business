package common

import "fmt"

const (
	LogLevelDanger  = 1
	LogLevelWarning = 2
	LogLevelNormal  = 3
)

func Log(level int, err error) {
	fmt.Println(err.Error())
}
