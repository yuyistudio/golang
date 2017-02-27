package benchmark

import (
	"testing"
)


type Person struct {
	name string
	age int
}

func BenchmarkB1(b *testing.B) {
	b.Logf("start")
	b.Error("log & fail")
	b.StopTimer()
	// do something irrelevant to timing
	b.StartTimer()
	// continue the measure of time
	b.FailNow()
	b.Logf("end")
}



