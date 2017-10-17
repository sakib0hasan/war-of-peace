package main

import (
	"os"
	"bufio"
	"fmt"
	"sync/atomic"
	"time"
	"bytes"
)

type ExampleTask []byte

var readOps int32 = 0

func (e ExampleTask) Execute() {
	key_slice := []byte(" ")
	key_slice3 := []byte("\n")
	atomic.AddInt32(&readOps, int32(bytes.Count(e,key_slice)+bytes.Count(e,key_slice3))+1)
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
