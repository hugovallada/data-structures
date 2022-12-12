package genericlist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var expected string = "ola"
	list := New[string]()
	list.Insert("ola")
	if list.data[0] != expected {
		t.Errorf("Expected %s, Got %s", expected, list.data[0])
	}
}

func TestInsertAll(t *testing.T) {
	var expected []string = []string{"ola", "mundo"}
	list := New[string]()
	list.InsertAll("ola", "mundo")
	if len(list.data) != len(expected) {
		t.Errorf("Expected a list of size %d, Got %d", len(list.data), len(expected))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != list.data[i] {
			t.Errorf("Expected %s, Got %s", expected[i], list.data[i])
		}
	}
}

func TestGet(t *testing.T) {
	var expected string = "aguia"
	list := New[string]()
	list.InsertAll("ola", "mundo", "feroz", "carro", "case", "aguia", "cachorro")
	value, err := list.Get(5)
	if err != nil {
		t.Error(err)
	}
	if value != expected {
		t.Errorf("Expected %s, Got %s", expected, value)
	}
}

func TestGet_ShouldReturnErrorWhenIndexIsOutOfBound(t *testing.T) {
	var expected string = "index is too high"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	_, err := list.Get(5)
	if err == nil {
		t.Errorf("Expected %s, got %s", expected, err)
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

func TestGet_ShouldReturnErrorWhenNegativeIndexIsGiven(t *testing.T) {
	var expected string = "index can't be negative"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	_, err := list.Get(-5)
	if err == nil {
		t.Errorf("Expected %s, got %s", expected, err)
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

func TestGetIndex(t *testing.T) {
	var expected = 5
	list := New[string]()
	list.InsertAll("ola", "mundo", "feroz", "carro", "case", "aguia", "cachorro")
	value, err := list.GetIndex("aguia")
	if err != nil {
		t.Error(err)
	}
	if value != expected {
		t.Errorf("Expected %d, Got %d", expected, value)
	}
}

func TestGetIndex_ShouldReturnErrorWhenValueNotFound(t *testing.T) {
	var expected = "value not found"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	_, err := list.GetIndex("aguia")
	if err == nil {
		t.Errorf("Expected %s, got %s", expected, err)
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

func TestRemove(t *testing.T) {
	var expected string = "mundo"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	removedValue, err := list.Remove(1)
	if err != nil {
		t.Errorf("Was not expecting an error, got %v", err)
	}
	if removedValue != expected {
		t.Errorf("Expected %s, Got %s", expected, removedValue)
	}
	if len(list.data) > 1 {
		t.Errorf("expected 1, but got %d\n", len(list.data))
	}
}

func TestRemove_ShouldReturnErrorWhenIndexOutOfBound(t *testing.T) {
	var expected string = "index is too high"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	_, err := list.Remove(5)
	if err == nil {
		t.Error("Was expecting an error, but got nothing")
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
	if len(list.data) < 2 {
		t.Errorf("expected 1, but got %d\n", len(list.data))
	}
}

func TestRemoveByValue(t *testing.T) {
	list := New[string]()
	list.InsertAll("ola", "mundo")
	err := list.RemoveByValue("ola")
	if err != nil {
		t.Errorf("Was not expecting an error, got %v", err)
	}
	if len(list.data) > 1 {
		t.Errorf("expected 1, but got %d\n", len(list.data))
	}
}

func TestRemoveByValue_ShouldReturnErrorWhenIndexOutOfBound(t *testing.T) {
	var expected string = "value not found"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	err := list.RemoveByValue("aguia")
	if err == nil {
		t.Error("Was expecting an error, but got nothing")
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
	if len(list.data) < 2 {
		t.Errorf("expected 1, but got %d\n", len(list.data))
	}
}

func TestUpdateValue(t *testing.T) {
	var newValue string = "aguia"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	err := list.UpdateValue("ola", newValue)
	if err != nil {
		t.Errorf("Was not expecting an error, got %v", err)
	}
	if list.data[0] != newValue {
		t.Errorf("Expected %s, Got %s", newValue, list.data[0])
	}
	if list.data[1] != "mundo" {
		t.Errorf("Expected %s, Got %s", "mundo", list.data[1])
	}

	if len(list.data) != 2 {
		t.Errorf("expected 2, but got %d\n", len(list.data))
	}
}

func TestUpdateValue_ShouldReturnAnErrorWhenValueNotFound(t *testing.T) {
	var newValue string = "aguia"
	var expected string = "value not found"
	list := New[string]()
	list.InsertAll("ola", "mundo")
	err := list.UpdateValue("casco", newValue)
	if err == nil {
		t.Error("Was expecting an error, but got nothing")
	}
	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
	if list.data[0] != "ola" {
		t.Errorf("Expected %s, Got %s", "ola", list.data[0])
	}
	if list.data[1] != "mundo" {
		t.Errorf("Expected %s, Got %s", "mundo", list.data[1])
	}
	if len(list.data) != 2 {
		t.Errorf("expected 2, but got %d\n", len(list.data))
	}
}
