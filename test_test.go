package main

import (
	"runtime"
	"testing"
)

func BenchmarkPokerCalculator(b *testing.B) {
	InitJSONToMap()
	a := make([]string, 7)
	a[0] = "144"
	a[1] = "131"
	a[2] = "121"
	a[3] = "122"
	a[4] = "91"
	a[5] = "24"
	a[6] = "132"
	for i := 0; i < b.N; i++ {
		PokerCalculator(a)
	}
}
func BenchmarkGoPokerCalculator(b *testing.B) {
	InitJSONToMap()
	configGOMAXPROCS := int(config["runtimeGOMAXPROCS"].(float64))
	runtime.GOMAXPROCS(configGOMAXPROCS)
	c := make(chan bool)
	a := make([]string, 7)
	a[0] = "144"
	a[1] = "131"
	a[2] = "121"
	a[3] = "122"
	a[4] = "91"
	a[5] = "24"
	a[6] = "132"
	for i := 0; i < b.N; i++ {
		go GoPokerCalculator(c, a)
	}
	<-c
}
