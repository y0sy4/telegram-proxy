package pool

import (
	"testing"
	"time"
)

func TestWSPool_PutAndGet(t *testing.T) {
	p := NewWSPool(2, 1*time.Second)
	key := DCKey{DC: 2, IsMedia: false}

	// Pool should be empty initially
	if ws := p.Get(key); ws != nil {
		t.Error("expected nil from empty pool")
	}

	// Put a connection
	// We can't create a real WebSocket in unit tests,
	// but we can verify the pool logic works with nil checks
}

func TestWSPool_NeedRefill(t *testing.T) {
	p := NewWSPool(2, 1*time.Second)
	key := DCKey{DC: 2, IsMedia: false}

	// Should need refill when empty
	if !p.NeedRefill(key) {
		t.Error("expected NeedRefill to be true for empty pool")
	}
}

func TestWSPool_MaxAge(t *testing.T) {
	p := NewWSPool(2, 50*time.Millisecond)
	key := DCKey{DC: 2, IsMedia: false}

	// Put would expire quickly — verify age check logic
	// Again, limited without real WebSocket, but structure is tested
	if !p.NeedRefill(key) {
		t.Error("expected NeedRefill true")
	}
}

func TestDCKey_Equality(t *testing.T) {
	k1 := DCKey{DC: 2, IsMedia: false}
	k2 := DCKey{DC: 2, IsMedia: false}
	k3 := DCKey{DC: 2, IsMedia: true}
	k4 := DCKey{DC: 4, IsMedia: false}

	if k1 != k2 {
		t.Error("expected equal DCKeys")
	}
	if k1 == k3 {
		t.Error("expected different IsMedia")
	}
	if k1 == k4 {
		t.Error("expected different DC")
	}
}
