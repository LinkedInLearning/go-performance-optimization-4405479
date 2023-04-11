package auth

import (
	"fmt"
	"testing"
)

var users Store

func init() {
	for i := 0; i < 10_000; i++ {
		u := User{
			Login: fmt.Sprintf("user-%04d", i),
			Token: fmt.Sprintf("tok-%04d", i),
		}
		users = append(users, u)
	}
}

func BenchmarkToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tok := users.Token("joe")
		if tok != "" {
			b.Fatal("found non-existing user")
		}
	}
}
