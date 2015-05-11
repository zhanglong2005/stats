package service

import (
	"github.com/bborbe/stats/entry"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

type service struct {
}

type Service interface {
	List() ([]entry.Entry, error)
	Create(entry *entry.Entry) (*entry.Entry, error)
}

func New(storage stats_entry_storage.Storage) *service {
	return new(service)
}

func (s *service) List() ([]entry.Entry, error) {
	return nil, nil
}

func (s *service) Create(entry *entry.Entry) (*entry.Entry, error) {
	return nil, nil
}
