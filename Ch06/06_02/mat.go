package mat

type Mat [][]float64

func (m Mat) Dim() (rows int, cols int) {
	return len(m), len(m[0])
}

func (m Mat) Sum() float64 {
	nRows, nCols := m.Dim()
	total := 0.0
	for r := 0; r < nRows; r++ {
		for c := 0; c < nCols; c++ {
			total += m[r][c]
		}
	}

	return total
}
