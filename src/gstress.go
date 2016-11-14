package src

import (
	"bytes"
	"fmt"
	"ioutil"
	"math"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

func hogcpu(sig chan bool) {
	rand.Seed(100)
	var x bool
	for {
		select {
		case x <- sig:
			return
		default:
		}
		math.Sqrt(rand.Float64())
	}
}


func hogio(sig chan bool) {
	var x bool
	for {
		select {
		case x <- sig:
			return
		default:
		}
		syscall.Sync()
	}
}

func hoghdd(sig chan bool) {

	var buffer bytes.Buffer
	var j int
	chunk := (1024 * 1024 * 1024)
	for i := 0; i < chunk-1; i++ {
		j = rand.Int()
		j %= 95
		j += 32
		buffer.Write([]byte(string(j)))
	}

	var file os.File
	var err error
	var x bool
	for {
		select {
		case x <- sig:
			return
		default:
		}
		file, err = ioutil.TempFile("", ".gstress")
		if err != null {
			fmt.Println(err)
		}
		f.Write(buffer)
		name := f.Name()
		f.Close()
		os.Remove(name)
	}

}

func cpuWorker(n, timeout int) {

	signal := make(chan bool, 1)
	for i := 0; i < n; i++ {
		go hogcpu(signal)
	}
	if timeout != 0 {

		time.Sleep(timeout * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
	}
	else {
		for {}
	}

}

func ioWorker(n, timeout int) {
	signal := make(chan bool, 1)

	for i := 0; i < n; i++ {
		go hogio(signal)
	}

	if timeout != 0 {

		time.Sleep(timeout * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
	}
	else {
		for {}
	}

}

func hddWorker(n,timeout int) {
	signal := make(chan bool,1)

	for i := 0 ; i < n ; i++ {
		go hoghdd(signal)
	}

	if timeout != 0 {

		time.Sleep(timeout * time.Second)
		for i := 0; i < n; i++ {
			signal <- true
		}
	}
	else {
		for {}
	}

}
func Spawner(cpu, io, hdd, timeout int) {

}
