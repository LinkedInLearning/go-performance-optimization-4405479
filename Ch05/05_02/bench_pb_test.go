package cart

import (
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go-optimize/Ch05/05_02/pb"
)

var (
	cartPB = pb.Cart{
		User:    "joe",
		Updated: timestamppb.New(time.Date(2023, time.January, 19, 14, 52, 30, 0, time.UTC)),
		Items: []*pb.Item{
			{Sku: "hammer19", Amount: 1, Price: 3.7},
			{Sku: "nail7", Amount: 100, Price: 0.01},
			{Sku: "glue6", Amount: 2, Price: 2.3},
		},
	}
)

func BenchmarkPB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&cartPB)
		if err != nil {
			b.Fatal(err)
		}
	}
}
