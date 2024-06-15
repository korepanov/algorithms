package main

import (
	"math"
	"testing"

	"github.com/shady831213/algorithms/sort"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	p := plot.New()
	p.Title.Text = "Merge sort time"
	p.X.Label.Text = "Array size"
	p.Y.Label.Text = "Time (ms)"

	var sizes = [...]int{1000000, 2000000, 3000000, 4000000, 5000000, 6000000, 7000000, 8000000, 9000000, 10000000}

	timeComplexityPts := make(plotter.XYs, len(sizes))
	experimentTimePts := make(plotter.XYs, len(sizes))

	for idx, size := range sizes {
		var b testing.B
		b.N = 10
		timeComplexityPts[idx].X = float64(size)
		timeComplexityPts[idx].Y = float64(size) * math.Log(float64(size))

		sort.BenchmarkSortWithSize(&b, sort.MergeSort, size)
		experimentTimePts[idx].X = float64(size)
		experimentTimePts[idx].Y = float64(int(float64(b.Elapsed().Milliseconds()) / float64(b.N)))
	}

	scaleKoef := experimentTimePts[len(sizes)-1].Y / timeComplexityPts[len(sizes)-1].Y

	for idx := range experimentTimePts {
		timeComplexityPts[idx].Y *= scaleKoef
	}

	err := plotutil.AddLinePoints(p,
		"Time compexity plot", timeComplexityPts,
		"Experiment time plot", experimentTimePts)

	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "result.png"); err != nil {
		panic(err)
	}

}
