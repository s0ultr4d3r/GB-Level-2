package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

var list [5]string

var (
	ErrIndex = errors.New("oversized")
)

func panicHandler() {
	if v := recover(); v != nil {
		buff := make([]byte, 1024)
		runtime.Stack(buff, false)
		moment := time.Now()
		fmt.Printf("Catched panic: %v, %s\n Time: %s\n", v, buff, moment)
	}
}
func overloadList() {

	for i := 0; i < 10; i++ {
		list[i] = "elem №: " + strconv.Itoa(i)
	}
}

func mixIt() {
	defer panicHandler()
	overloadList()
}

func createMilFiles() {
	for i := 0; i < int(math.Pow(10, 6)); i++ {
		emptyFile, err := os.Create("emptyFile" + strconv.Itoa(i))
		if err != nil {
			if v := recover(); v != nil {
				buff := make([]byte, 1024)
				runtime.Stack(buff, false)
				moment := time.Now()
				fmt.Printf("Limit files reached: %v, %s\n Time: %s\n", v, buff, moment)
			}
		}
		defer emptyFile.Close()
	}
}

func main() {
	mixIt()
	fmt.Println(list)
	fmt.Println("panic overcoming")

	createMilFiles()
}
