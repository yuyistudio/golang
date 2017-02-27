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
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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
	loop2(count)
}
func loop2(count int) {
	loop(count / 10)
	loop11(count, false)
	for i := 0; i < count; i++ {
		tmp := i
		i += 1
		i = tmp
	}
}
func delay(seconds int64) {
	defer func(){done<-true}()
	fmt.Println("delay", seconds)
	loop(int(seconds * 1e8))
	//time.Sleep(time.Duration(seconds) * time.Second)
	loop11(int(seconds * 5e7), true)
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
	for i := 0; i < count; i++ {
		fmt.Println("waiting", i)
		<-done
	}
}
