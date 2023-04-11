package user

import (
	"fmt"
	"testing"
)

var (
	sdb        SliceCache
	mdb        MapCache
	n          = 10
	benchLogin = "no such user"
)

func init() {
	mdb = make(MapCache)
	for i := 0; i < n; i++ {
		u := User{Login: fmt.Sprintf("u-%04d", i)}
		sdb = append(sdb, u)
		mdb[u.Login] = u
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := sdb.Find(benchLogin)
		if ok {
			b.Fatal("found non-existing login")
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := mdb.Find(benchLogin)
		if ok {
			b.Fatal("found non-existing login")
		}
	}
}
