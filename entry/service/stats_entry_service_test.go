package service

import (
	"testing"

	. "github.com/bborbe/assert"
	stats_entry "github.com/bborbe/stats/entry"
	stats_entry_storage "github.com/bborbe/stats/entry/storage"
)

func TestNewHandlerImplementsService(t *testing.T) {
	r := New(nil)
	var i (*Service) = nil
	err := AssertThat(r, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	var err error
	storage := stats_entry_storage.New("/tmp/stats_test.db", false)
	service := New(storage)
	if err = storage.Truncate(); err != nil {
		t.Fatal(err)
	}
	entries, err := storage.FindEntries()
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(entries), Is(0)); err != nil {
		t.Fatal(err)
	}
	entry, err := service.Create(&stats_entry.Entry{Value: 1337})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(entry, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(entry.Value, Is(1337)); err != nil {
		t.Fatal(err)
	}
	entries, err = storage.FindEntries()
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(entries), Is(1)); err != nil {
		t.Fatal(err)
	}
}

func TestList(t *testing.T) {
	var err error
	storage := stats_entry_storage.New("/tmp/stats_test.db", false)
	service := New(storage)
	if err = storage.Truncate(); err != nil {
		t.Fatal(err)
	}
	entries, err := service.List()
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(entries), Is(0)); err != nil {
		t.Fatal(err)
	}
	err = storage.CreateEntry(&stats_entry.Entry{Value: 1337})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	entries, err = service.List()
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(entries), Is(1)); err != nil {
		t.Fatal(err)
	}
}
