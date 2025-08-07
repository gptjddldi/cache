package sieve_test

import (
	"sieve"
	"testing"
)

func TestSieve_PutAndGet(t *testing.T) {
	cache := sieve.NewSieve(3)

	// 기본 Put/Get 테스트
	cache.Put(1, 100)
	cache.Put(2, 200)

	if got := cache.Get(1); got != 100 {
		t.Errorf("Expected 100, got %d", got)
	}
	if got := cache.Get(2); got != 200 {
		t.Errorf("Expected 200, got %d", got)
	}
	if got := cache.Get(3); got != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", got)
	}
}

func TestSieve_UpdateExisting(t *testing.T) {
	cache := sieve.NewSieve(3)

	cache.Put(1, 100)
	cache.Put(1, 150) // 기존 키 업데이트

	if got := cache.Get(1); got != 150 {
		t.Errorf("Expected 150, got %d", got)
	}
}

func TestSieve_CapacityLimit(t *testing.T) {
	cache := sieve.NewSieve(2)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)

	// 첫 번째 항목이 evict되어야 함
	if got := cache.Get(1); got != -1 {
		t.Errorf("Expected -1 for evicted key, got %d", got)
	}
	if got := cache.Get(2); got != 200 {
		t.Errorf("Expected 200, got %d", got)
	}
	if got := cache.Get(3); got != 300 {
		t.Errorf("Expected 300, got %d", got)
	}
}

func TestSieve_EvictionPolicy(t *testing.T) {
	cache := sieve.NewSieve(3)

	// 3개 항목 추가
	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)

	// 1번 키에 접근하여 visited = true로 설정
	cache.Get(1)

	// 4번째 항목 추가 - visited = false인 항목이 evict되어야 함
	cache.Put(4, 400)

	// 2번이 evict되어야 함 (visited = false)
	if got := cache.Get(2); got != -1 {
		t.Errorf("Expected -1 for evicted key 2, got %d", got)
	}
	if got := cache.Get(1); got != 100 {
		t.Errorf("Expected 100 for visited key, got %d", got)
	}
	if got := cache.Get(3); got != 300 {
		t.Errorf("Expected 300, got %d", got)
	}
	if got := cache.Get(4); got != 400 {
		t.Errorf("Expected 400, got %d", got)
	}
}

func TestSieve_ComplexEviction(t *testing.T) {
	cache := sieve.NewSieve(4)

	// 4개 항목 추가
	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)
	cache.Put(4, 400)

	// 모든 항목에 접근하여 visited = true로 설정
	cache.Get(1)
	cache.Get(2)
	cache.Get(3)
	cache.Get(4)

	// 5번째 항목 추가 - 모든 항목이 visited = true이므로
	// hand가 Back()에서 시작하여 첫 번째로 visited = false가 되는 항목을 찾아 evict
	cache.Put(5, 500)

	if got := cache.Get(4); got != 400 {
		t.Errorf("Expected 400 for evicted key 4, got %d", got)
	}
	// 1번이 evict되어야 함 (hand가 Back()에서 시작하므로)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Expected -1, got %d", got)
	}
	if got := cache.Get(2); got != 200 {
		t.Errorf("Expected 200, got %d", got)
	}
	if got := cache.Get(3); got != 300 {
		t.Errorf("Expected 300, got %d", got)
	}
	if got := cache.Get(5); got != 500 {
		t.Errorf("Expected 500, got %d", got)
	}
}

func TestSieve_HandMovement(t *testing.T) {
	cache := sieve.NewSieve(3)

	// 3개 항목 추가
	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)

	// 1번에 접근
	cache.Get(1)

	// 4번째 항목 추가
	cache.Put(4, 400)

	// hand가 2번 위치로 이동했는지 확인
	// 2번이 evict되어야 함
	if got := cache.Get(2); got != -1 {
		t.Errorf("Expected -1 for evicted key 2, got %d", got)
	}
	if got := cache.Get(1); got != 100 {
		t.Errorf("Expected 100, got %d", got)
	}
	if got := cache.Get(3); got != 300 {
		t.Errorf("Expected 300, got %d", got)
	}
	if got := cache.Get(4); got != 400 {
		t.Errorf("Expected 400, got %d", got)
	}
}

func TestSieve_EmptyCache(t *testing.T) {
	cache := sieve.NewSieve(3)

	// 빈 캐시에서 Get
	if got := cache.Get(1); got != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", got)
	}
	if got := cache.Get(2); got != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", got)
	}
}

func TestSieve_SingleItem(t *testing.T) {
	cache := sieve.NewSieve(1)

	cache.Put(1, 100)
	if got := cache.Get(1); got != 100 {
		t.Errorf("Expected 100, got %d", got)
	}

	// 두 번째 항목 추가
	cache.Put(2, 200)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Expected -1 for evicted key, got %d", got)
	}
	if got := cache.Get(2); got != 200 {
		t.Errorf("Expected 200, got %d", got)
	}
}

func TestSieve_RepeatedAccess(t *testing.T) {
	cache := sieve.NewSieve(2)

	cache.Put(1, 100)
	cache.Put(2, 200)

	// 같은 키에 여러 번 접근
	cache.Get(1)
	cache.Get(1)
	cache.Get(2)

	// 3번째 항목 추가
	cache.Put(3, 300)

	// 1번이 evict되어야 함
	if got := cache.Get(1); got != -1 {
		t.Errorf("Expected -1, got %d", got)
	}
	if got := cache.Get(2); got != 200 {
		t.Errorf("Expected 200 for evicted key 2, got %d", got)
	}
	if got := cache.Get(3); got != 300 {
		t.Errorf("Expected 300, got %d", got)
	}
}
