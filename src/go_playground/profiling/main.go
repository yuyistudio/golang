/*
profile一秒钟采样100次,这一秒钟是实际运行的累加时间,不包括sleep的时间.
所以依靠sleep去模拟延时是无法得到采样数据的.应该使用for循环去模拟延时.
*/
package main

import (
	"log"
	"os"
	"flag"
	"runtime/pprof"
	//"time"
	"fmt"
	"sync"
	//"runtime"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var print = make(chan bool)
var dummies []*Dummy
var lock sync.Mutex

type Dummy struct {
	age int
	name string
}

var done = make(chan bool)
func loop(count int) {
	for i := 0; i < count; i++ {
		tmp := i
		i += 1
		i = tmp
	}
}

func loop11(count int, recursive bool) {
	if recursive {
		loop12(count)
	} else {
		loop(count / 100)
	}
}
func loop12(count int) {
	for _, dummpy := range loop2(count) {
		if dummpy.age < 1 {
			fmt.Print("error")
		}
	}
}
func AllocateDummy(i int) *Dummy {
	d := new(Dummy)
	d.age = 13
	d.name = fmt.Sprintf("Jecco is %v years old", i)
	return d
}
func loop2(count int) []*Dummy {
	loop(count / 10)
	loop11(count, false)
	for i := 0; i < count; i++ {
		d := AllocateDummy(i)
		tmp := i
		i += 1
		i = tmp
		d.age = tmp
		lock.Lock()
		dummies = append(dummies, d)
		lock.Unlock()
	}
	lock.Lock()
	for _, dummy := range dummies {
		dummy.age = 11
	}
	lock.Unlock()
	return dummies
}
func delay(seconds int64) {
	defer func(){done<-true}()
	fmt.Println("delay", seconds)
	loop(int(seconds * 1e5))
	//time.Sleep(time.Duration(seconds) * time.Second)
	loop11(int(seconds * 5e5), true)
	fmt.Println("delay done", seconds)
}
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
		    log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	count := 8
	for i := 0; i < count; i++ {
		go delay(int64(i))
	}
	for _, dummpy := range loop2(count) {
		if dummpy.age < 1 {
			fmt.Print("error")
		}
	}
	for i := 0; i < count; i++ {
		fmt.Println("waiting", i)
		<-done
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
