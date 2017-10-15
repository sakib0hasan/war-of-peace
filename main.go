package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

type ExampleTask []byte

var readOps int32 = 0

func (e ExampleTask) Execute() {
	line := string(e)
	line = strings.TrimSpace(line)
	spaces := strings.Count(line, " ")
	line_breaks := strings.Count(line, "\n")
	atomic.AddInt32(&readOps, int32(spaces)+1+int32(line_breaks))
}

func main() {

	start := time.Now()

	file, err := os.Open("wap.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buffer := bufio.NewReader(file)

	bytes := make([]byte, 100000)
	pool := NewPool(1000)
	for {
		v, _ := buffer.Read(bytes)
		if v != 0 {
			pool.Exec(ExampleTask(bytes))
		} else {
			break
		}
	}

	pool.Close()
	pool.Wait()
	fmt.Println(readOps)
	fmt.Println(time.Since(start))
}
