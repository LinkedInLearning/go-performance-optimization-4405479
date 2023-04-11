package mat

import "testing"

var (
	mat Mat
)

const (
	nRows = 1000
	nCols = 20
)

func init() {
	for r := 0; r < nRows; r++ {
		row := make([]float64, nCols)
		for c := 0; c < nCols; c++ {
			row[c] = float64(c)
		}
		mat = append(mat, row)
	}
}

func BenchmarkSumRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := mat.SumRows()
		if size := len(v); size != nRows {
			b.Fatal(size)
		}
	}
}
