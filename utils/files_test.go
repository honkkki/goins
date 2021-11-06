package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	path := "/home/honki/nacos"
	assert.Equal(t, true, Exist(path))

	path = ""
	assert.Equal(t, false, Exist(path))

	path = "/etc///"
	t.Log(Exist(path))
}

func BenchmarkExist(b *testing.B) {
	path := "/etc/apt"
	for i := 0; i < b.N; i++ {
		Exist(path)
	}
}
