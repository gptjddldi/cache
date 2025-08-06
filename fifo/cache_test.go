package fifo_test

import (
	"fifo"
	"testing"
)

func TestFifo(t *testing.T) {
	cache := fifo.NewCache(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)

	if cache.Get(1) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(1))
	}

	if cache.Get(2) != 2 {
		t.Errorf("Expected 2, got %d", cache.Get(2))
	}

	if cache.Get(3) != 3 {
		t.Errorf("Expected 3, got %d", cache.Get(3))
	}

	if cache.Get(4) != 4 {
		t.Errorf("Expected 4, got %d", cache.Get(4))
	}
}
