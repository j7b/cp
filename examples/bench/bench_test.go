package main

import "testing"

var space = simpleTerrainCircles_1000()

func BenchmarkStep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		space.Step(1.0 / 60.0)
	}
}
