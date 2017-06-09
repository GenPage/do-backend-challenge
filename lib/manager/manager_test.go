package manager

import (
	"testing"

	"github.com/genpage/do-backend-challenge/lib/model"
)

func TestIndex(t *testing.T) {
	testData := "INDEX|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err != nil {
		t.Error(err)
	}
}

func TestIndexAlreadyExists(t *testing.T) {
	testData := "INDEX|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err != nil {
		t.Error(err)
	}

	err = m.Index(p)

	if err == nil {
		t.Error("Error not thrown")
	}
}

func TestIndexMissingDep(t *testing.T) {
	testData := "INDEX|test2|test"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err == nil {
		t.Error("Error not thrown")
	}
}

func TestQuery(t *testing.T) {
	testData := "QUERY|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err != nil {
		t.Error(err)
	}

	err = m.Query(p)

	if err != nil {
		t.Error(err)
	}
}

func TestQueryGone(t *testing.T) {
	testData := "QUERY|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Query(p)

	if err == nil {
		t.Error("Error not thrown")
	}
}

func TestRemove(t *testing.T) {
	testData := "REMOVE|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err != nil {
		t.Error(err)
	}

	err = m.Remove(p)

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveAlreadyGone(t *testing.T) {
	testData := "REMOVE|test|"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Remove(p)

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveDepedencyCheck(t *testing.T) {
	testData := "REMOVE|test|"
	testData2 := "INDEX|test2|test"

	p, err := model.PackageFromString(testData)
	if err != nil {
		t.Error("ERROR")
	}

	p2, err := model.PackageFromString(testData2)
	if err != nil {
		t.Error("ERROR")
	}

	m := NewManager()

	err = m.Index(p)

	if err != nil {
		t.Error(err)
	}

	err = m.Index(p2)

	if err != nil {
		t.Error(err)
	}

	err = m.Remove(p)

	if err == nil {
		t.Error("Error not thrown")
	}
}
