package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)
type ExampleTask string
var readOps int32 = 0

func (e ExampleTask) Execute() {
	parts := strings.Count(string(e), " ")
	atomic.AddInt32(&readOps, int32(parts+1))
}

func main() {
	start := time.Now()
	file, err := os.Open("wap.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	info, _ := file.Stat()

	// calculate the bytes size
	var size int64 = info.Size()
	bytes := make([]byte, size)
	byte_size := int(info.Size())
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	pool := NewPool(10)
	for i := 0; i < byte_size; i += 90000 {
		ba := bytes[i:min(i+90000, byte_size)]
		pool.Exec(ExampleTask(ba))
	}
	pool.Close()
	pool.Wait()
	fmt.Println(readOps)
	fmt.Println(time.Since(start))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}