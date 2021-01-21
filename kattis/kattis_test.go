package kattis_test

import (
	"os"
	"testing"

	"github.com/norbs57/exercises/kattis"
	"github.com/stretchr/testify/assert"
)

func TestKattis(t *testing.T) {
	tsps := []string{"10", "52", "280", "1000", "1173"}
	costs := []int{331, 8708, 3132, 24556237, 72059}
	for i, s := range tsps {
		fName := "../data/tsp" + s + ".txt"
		f, err := os.Open(fName)
		if err == nil {
			os.Stdin = f
			defer f.Close()
		}
		_, cost := kattis.Tsp()
		assert.Equal(t, cost, costs[i])
	}
}
