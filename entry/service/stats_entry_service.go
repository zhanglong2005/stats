package service

import (
	"time"

	"github.com/bborbe/stats/entry"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

type service struct {
	storage stats_entry_storage.Storage
}

type Service interface {
	List() ([]entry.Entry, error)
	ListLimited(limit int) ([]entry.Entry, error)
	Create(entry *entry.Entry) (*entry.Entry, error)
}

func New(storage stats_entry_storage.Storage) *service {
	s := new(service)
	s.storage = storage
	return s
}

func (s *service) List() ([]entry.Entry, error) {
	return s.storage.FindEntries()
}

func (s *service) ListLimited(limit int) ([]entry.Entry, error) {
	return s.storage.FindLatestEntries(limit)
}

func (s *service) Create(e *entry.Entry) (*entry.Entry, error) {
	timestamp := time.Now().UnixNano()
	e = &entry.Entry{Value: e.Value, Timestamp: timestamp}
	err := s.storage.CreateEntry(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
