package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCookie(t *testing.T) {
	ret, err := GetCookie()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, r)
	t.Log(ret)
}

func BenchmarkGetCookie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetCookie()
		if err != nil {
			b.Fatal(err)
		}
	}
}