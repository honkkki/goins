package utils

import "testing"

func BenchmarkGetCookie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetCookie()
		if err != nil {
			b.Fatal(err)
		}
	}
}