package state

import "log/slog"

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
}

func NewInMemoryStorage() *InMemoryStorage {
	slog.Info("in-memory storage created")
	return &InMemoryStorage{data: make(map[int64]*UserState)}
}

func (s *InMemoryStorage) Get(userID int64) (*UserState, error) {
	return s.data[userID], nil
}

func (s *InMemoryStorage) Set(userID int64, state *UserState) error {
	s.data[userID] = state
	return nil
}

func (s *InMemoryStorage) Delete(userID int64) error {
	delete(s.data, userID)
	return nil
}
