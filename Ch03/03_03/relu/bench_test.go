package relu

import "testing"

var values []int

const n = 10_000

func init() {
	values = make([]int, n)
	for i := range values {
		if i%2 == 0 {
			values[i] = i
		} else {
			values[i] = -i
		}
	}
}

func BenchmarkRelu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := Relu(values)
		if len(v) != n {
			b.Fatal(len(v))
		}
	}
}
