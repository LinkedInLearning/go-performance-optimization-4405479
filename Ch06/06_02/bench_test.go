package mat

import "testing"

var (
	mat          Mat
	nRows, nCols = 17, 9231
)

func init() {
	mat = make([][]float64, nRows)
	for r := 0; r < nRows; r++ {
		row := make([]float64, nCols)
		for c := 0; c < nCols; c++ {
			row[c] = 1
		}
		mat[r] = row
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := mat.Sum()
		if int(s) != nRows*nCols {
			b.Fatal(s)
		}
	}
}
