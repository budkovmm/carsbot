package state

import (
	"sync"
	"testing"
)

func TestInMemoryStorage_SetGetDelete(t *testing.T) {
	storage := NewInMemoryStorage()
	userID := int64(42)
	state := &UserState{Step: 1, SellerName: "Ivan"}

	// Set
	if err := storage.Set(userID, state); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Get
	got, err := storage.Get(userID)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got == nil || got.SellerName != "Ivan" {
		t.Errorf("Get returned wrong state: %+v", got)
	}

	// Delete
	if err := storage.Delete(userID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	got, _ = storage.Get(userID)
	if got != nil {
		t.Errorf("State not deleted: %+v", got)
	}
}

func TestInMemoryStorage_ConcurrentAccess(t *testing.T) {
	storage := NewInMemoryStorage()
	wg := sync.WaitGroup{}
	userID := int64(42)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			state := &UserState{Step: i}
			storage.Set(userID, state)
			got, _ := storage.Get(userID)
			if got == nil {
				t.Errorf("Get returned nil in goroutine %d", i)
			}
		}(i)
	}
	wg.Wait()
}
