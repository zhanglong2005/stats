package storage

import (
	"github.com/bborbe/stats/entry"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	FindEntrys() ([]entry.Entry, error)
	CreateEntry(entry *entry.Entry) error
	GetEntry(id int) (*entry.Entry, error)
	DeleteEntry(id int) (*entry.Entry, error)
}

type storage struct {
	db      *gorm.DB
	dbpath  string
	logmode bool
}

func New(dbpath string, logmode bool) *storage {
	s := new(storage)
	s.dbpath = dbpath
	s.logmode = logmode
	return s
}

func (s *storage) Truncate() error {
	db, err := s.getDb()
	if err != nil {
		return err
	}
	err = db.DropTableIfExists(&entry.Entry{}).Error
	if err != nil {
		return err
	}
	err = db.CreateTable(&entry.Entry{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) getDb() (*gorm.DB, error) {
	if s.db == nil {
		db, err := gorm.Open("sqlite3", s.dbpath)
		if err != nil {
			return nil, err
		}
		db.SingularTable(true)
		db.LogMode(s.logmode)
		db.AutoMigrate(&entry.Entry{})
		s.db = &db
	}
	return s.db, nil
}

func (s *storage) FindEntrys() ([]entry.Entry, error) {
	db, err := s.getDb()
	if err != nil {
		return nil, err
	}
	entrys := &[]entry.Entry{}
	query := db.Find(entrys)
	return *entrys, query.Error
}

func (s *storage) CreateEntry(entry *entry.Entry) error {
	db, err := s.getDb()
	if err != nil {
		return err
	}
	query := db.Create(entry)
	return query.Error
}

func (s *storage) GetEntry(id int) (*entry.Entry, error) {
	db, err := s.getDb()
	if err != nil {
		return nil, err
	}
	entry := &entry.Entry{}
	query := db.First(entry, id)
	if query.Error != nil {
		return nil, err
	}
	return entry, nil
}

func (s *storage) DeleteEntry(id int) (*entry.Entry, error) {
	db, err := s.getDb()
	if err != nil {
		return nil, err
	}
	entry := &entry.Entry{}
	query := db.First(entry, id)
	if query.Error != nil {
		return nil, err
	}
	query = db.Delete(entry)
	if query.Error != nil {
		return nil, err
	}
	return entry, nil
}
