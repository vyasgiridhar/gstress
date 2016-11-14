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

func hogcpu(sig chan bool) {
	rand.Seed(100)
	for {
		select {
		case <-sig:
			return
		default:
		}
		fmt.Println("Calculating")
		_ = math.Sqrt(rand.Float64())
	}
}

func hogio(sig chan bool) {
	for {
		select {
		case <-sig:
			return
		default:
		}
		fmt.Println("Syncing")
		syscall.Sync()
	}
}

func hoghdd(sig chan bool) {

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
		select {
		case <-sig:
			return
		default:
		}
		fmt.Println("Writing to file")
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

	signal := make(chan bool, 1)

	defer wait.Done()

	for i := 0; i < n; i++ {
		go hogcpu(signal)
	}
	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
	} else {
		for {
		}
	}

}

func ioWorker(n, timeout int, wait *sync.WaitGroup) {
	signal := make(chan bool, 1)
	defer wait.Done()
	for i := 0; i < n; i++ {
		go hogio(signal)
	}

	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
	} else {
		for {
		}
	}

}

func hddWorker(n, timeout int, wait *sync.WaitGroup) {
	signal := make(chan bool, 1)
	defer wait.Done()
	for i := 0; i < n; i++ {
		go hoghdd(signal)
	}

	if timeout != 0 {

		time.Sleep(time.Duration(timeout) * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
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
