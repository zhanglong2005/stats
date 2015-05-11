package service

import (
	"github.com/bborbe/stats/entry"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

type service struct {
	storage stats_entry_storage.Storage
}

type Service interface {
	List() ([]entry.Entry, error)
	Create(entry *entry.Entry) (*entry.Entry, error)
}

func New(storage stats_entry_storage.Storage) *service {
	s := new(service)
	s.storage = storage
	return s
}

func (s *service) List() ([]entry.Entry, error) {
	return s.storage.FindEntrys()
}

func (s *service) Create(entry *entry.Entry) (*entry.Entry, error) {
	err := s.storage.CreateEntry(entry)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
