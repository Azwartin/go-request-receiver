package server

import (
	"strconv"
	"testing"

	"github.com/azwartin/go-request-receiver/tasks"
)

func TestCreation(t *testing.T) {
	storage := GetStorage()
	if storage != GetStorage() {
		t.Error("Storage is not singleton")
	}
}

func TestStore(t *testing.T) {
	storage := GetStorage()
	tests := map[string]interface{}{
		"a": "a",
		"b": 120,
		"c": tasks.URLFetchResult{},
	}

	for key, expect := range tests {
		storage.Store(key, expect)
		_, ok := storage.Load(key)
		if !ok {
			t.Errorf("Expect %v after store, but not found", expect)
		}
	}
}

func TestLoad(t *testing.T) {
	storage := GetStorage()
	key := "key"
	task := tasks.URLFetchTask{
		ID: "test",
	}

	storage.Store(key, task)
	got, ok := storage.Load(key)
	if !ok {
		t.Errorf("Expect %v after store, but not found", task)
	}

	if got.(tasks.URLFetchTask).ID != task.ID {
		t.Errorf("Expect id %v after store, but got %v", task.ID, got.(tasks.URLFetchTask).ID)
	}
}

type RangeTest struct {
	limit  uint
	offset uint
	length uint
}

func TestLoadRange(t *testing.T) {
	if instance != nil && len(instance.items) > 0 {
		instance.items = make(map[string]interface{})
	}

	storage := GetStorage()
	for i := 0; i < 10; i++ {
		id := strconv.Itoa(i)
		storage.Store(id, tasks.URLFetchTask{ID: id})
	}

	tests := []*RangeTest{
		&RangeTest{
			3, 0, 3,
		},
		&RangeTest{
			15, 0, 10,
		},
		&RangeTest{
			7, 3, 7,
		},
		&RangeTest{
			1, 5, 1,
		},
		&RangeTest{
			9, 9, 1,
		},
		&RangeTest{
			15, 11, 0,
		},
	}

	for _, test := range tests {
		items := storage.LoadRange(test.limit, test.offset)
		if len(items) != int(test.length) {
			t.Errorf("Expect len %v, but got %v", test.length, len(items))
		}

		if len(items) > 0 && items[0].(tasks.URLFetchTask).ID != strconv.Itoa(int(test.offset)) {
			t.Errorf("Expect id %v, but got %v", test.offset, items[0].(tasks.URLFetchTask).ID)
		}
	}
}

func TestDelete(t *testing.T) {
	storage := GetStorage()
	key := "key"
	task := tasks.URLFetchTask{
		ID: "test",
	}

	storage.Store(key, task)
	storage.Delete(key)
	if _, ok := storage.Load(key); ok {
		t.Error("Expect not found after delete, but found")
	}
}
