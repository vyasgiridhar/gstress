package src

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"sync"
	"syscall"
	"time"
)

func hogcpu() {
	rand.Seed(1000)
	for {
		_ = math.Sqrt(rand.Float64())
	}
}

func hogio() {
	for {
		syscall.Sync()
	}
}

func hoghdd() {

	var buffer bytes.Buffer
	var j int
	chunk := (1024 * 1024)
	for i := 0; i < chunk-1; i++ {
		j = rand.Int()
		j %= 95
		j += 32
		buffer.Write([]byte(string(j)))
	}

	var file *os.File
	var err error

	for {
		file, err = ioutil.TempFile("", ".gstress")
		if err != nil {
			fmt.Println(err)
		}
		file.Write(buffer.Bytes())
		name := file.Name()
		file.Close()
		os.Remove(name)
	}

}

func cpuWorker(n, timeout int, wait *sync.WaitGroup) {

	defer wait.Done()

	for i := 0; i < n; i++ {
		go hogcpu()
	}
	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
		return
	} else {
		for {
		}
	}

}

func ioWorker(n, timeout int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < n; i++ {
		go hogio()
	}

	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
	} else {
		for {
		}
	}

}

func hddWorker(n, timeout int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < n; i++ {
		go hoghdd()
	}

	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
	} else {
		for {
		}
	}

}
func Spawner(cpu, io, hdd, timeout int) {

	var wg sync.WaitGroup
	wg.Add(3)
	go cpuWorker(cpu, timeout, &wg)
	go hddWorker(hdd, timeout, &wg)
	go ioWorker(io, timeout, &wg)
	wg.Wait()
}
