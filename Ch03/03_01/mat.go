package mat

type Mat [][]float64

func (m Mat) SumRows() []float64 {
	var sums []float64
	for _, row := range m {
		sum := 0.0
		for _, v := range row {
			sum += v
		}
		sums = append(sums, sum)
	}

	return sums
}
