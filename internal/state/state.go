package state

import (
	"log/slog"
	"sync"
)

type UserState struct {
	Step       int
	SellerName string
	BuyerName  string
	VIN        string
	BrandModel string
	Year       string
	Color      string
	Price      string
	Date       string
	City       string
}

type StateStorage interface {
	Get(userID int64) (*UserState, error)
	Set(userID int64, state *UserState) error
	Delete(userID int64) error
}

type InMemoryStorage struct {
	data map[int64]*UserState
	mu   sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	slog.Info("in-memory storage created")
	return &InMemoryStorage{data: make(map[int64]*UserState)}
}

func (s *InMemoryStorage) Get(userID int64) (*UserState, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[userID], nil
}

func (s *InMemoryStorage) Set(userID int64, state *UserState) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[userID] = state
	return nil
}

func (s *InMemoryStorage) Delete(userID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, userID)
	return nil
}
