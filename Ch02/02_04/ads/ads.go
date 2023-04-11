package ads

type Impression struct {
	Price float64
}

func Total(imps []Impression) float64 {
	total := 0.0
	for _, i := range imps {
		total += i.Price
	}

	return total
}
