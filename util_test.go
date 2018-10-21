package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenValue(t *testing.T) {
	initRandSeed()
	v1 := getUniqueValue()
	v2 := getUniqueValue()

	if v1 == v2 {
		t.Errorf("Unique values %q and %q are equal", v1, v2)
	}

}

func TestMapCopy(t *testing.T) {
	src := mapType{
		"a": 1,
		"b": 2,
	}

	dest := make(mapType, len(src))
	copyMap(dest, src)

	if len(dest) != len(src) {
		t.Errorf("Map dest has unexpected length of %d", len(dest))
	}
	for k, v := range src {
		if _, ok := dest[k]; !ok {
			t.Errorf("Key %q does not exist in dest", k)
		}

		if v != dest[k] {
			t.Errorf("Expected value %q is not equal actual %q in dest", v, dest[k])
		}

	}

}

func TestCopyMethod(t *testing.T) {
	n := 100
	rand.Seed(time.Now().UnixNano())
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n)
	}

	data_copy := make([]int, len(data))
	copy(data_copy, data)
	for i, v := range data_copy {
		if v != data[i] {
			t.Errorf("Actual value %d at index %d is not equal expected %d", v, i, data[i])
		}
	}
}
