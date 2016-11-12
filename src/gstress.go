package src

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"syscall"
	"unsafe"
)

func hogcpu() {
	rand.Seed(100)
	for {
		math.Sqrt(rand.Float64())
	}
}

func hogvm() {
	for {
		var buffer bytes.Buffer
		x := 'A'
		for i := 0; i < 1024; i++ {
			buffer.WriteString("A")
		}
		time.sleep(1)
		for i := 0; i < 1024; i++ {
			{
				b, err := buffer.ReadByte()
				if err != nil {
					if string(b) != "A" {
						fmt.Println("Memory corruption at %v", unsafe.Pointer(s[i]))
					}
				}
			}
		}
	}
}

func hogio() {
	for {
		syscall.Sync()
	}
}

func hoghdd(size int64) {
	var buffer bytes.Buffer
	var j int
	chunk := (1024 * 1024)
	for i := 0; i < chunk-1; i++ {
		j = rand.Int()
		j %= 95
		j += 32
		buffer.Write([]byte(string(j)))
	}

}
