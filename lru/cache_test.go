package lru_test

import (
	"lru"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := lru.Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Errorf("Get(1) = %d; want 1", cache.Get(1))
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Errorf("Get(2) = %d; want -1", cache.Get(2))
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Errorf("Get(1) = %d; want -1", cache.Get(1))
	}
	if cache.Get(3) != 3 {
		t.Errorf("Get(3) = %d; want 3", cache.Get(3))
	}
	if cache.Get(4) != 4 {
		t.Errorf("Get(4) = %d; want 4", cache.Get(4))
	}
}
