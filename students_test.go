package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var p0 = Person{firstName: "Gavin", lastName: "Hill", birthDay: time.Date(1970, 6, 23, 0, 0, 0, 0, time.UTC)}
var p1 = Person{firstName: "Gavin", lastName: "Welch", birthDay: time.Date(1970, 6, 23, 0, 0, 0, 0, time.UTC)}
var p2 = Person{firstName: "Dan", lastName: "Dyer", birthDay: time.Date(1970, 6, 23, 0, 0, 0, 0, time.UTC)}
var p3 = Person{firstName: "Dave", lastName: "Rees", birthDay: time.Date(1988, 3, 24, 0, 0, 0, 0, time.UTC)}
var p4 = Person{firstName: "Dana", lastName: "Rees", birthDay: time.Date(1991, 8, 10, 0, 0, 0, 0, time.UTC)}
var people = People{p0, p1, p2, p3, p4}

func TestLen(t *testing.T) {
	res := people.Len()
	expected := 5
	
	if res != expected {
		t.Errorf("expected len: %d, got %d", expected, res)
	}
}

func TestLessName(t *testing.T) {
	res := people.Less(1, 0)
	expected := false

	if res != expected {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestLessBirth(t *testing.T) {
	res := people.Less(1, 2)
	expected := false

	if res != expected {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestLessDiff(t *testing.T) {
	res := people.Less(3, 4)
	expected := false

	if res != expected {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestSameBirth(t *testing.T) {
	res := people.Less(2, 3)
	expected := false

	if res != expected {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestSwap(t *testing.T) {
	people.Swap(1, 2)
	expected := "Dyer"
	res := people[1].lastName

	if res != expected {
		t.Errorf("expected: %s, got %s", expected, res)
	}
}

// TESTING MATRIX

func TestRows(t *testing.T) {
	m := `1 2
		  	3 4 
		  	5 6`
	res, _ := New(m)
	r := res.Rows()
	
	if r == nil {
		t.Errorf("expected correct rows returned")
	}
}

func TestCols(t *testing.T) {
	m := `1 2
		  	3 4 
		  	5 6`
	res, _ := New(m)
	cols := res.Cols()

	if cols == nil {
		t.Errorf("expected correct cols")
	}
}

func TestIncludeLetter(t *testing.T) {
	m := `1 2
		  	3 L 
		  	5 6`
	_, res := New(m)
	
	if res == nil {
		t.Errorf("expected numbers but got letter")
	}
}

func TestRowsSameLength(t *testing.T) {
	m := `1 2
		  	3 4 
		  	5 6`
	_, res := New(m)
	expected := true
	
	if res != nil {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestRowsDiffLength(t *testing.T) {
	m := `1
		  	3 4 
		  	5 6`
	_, err := New(m)
	
	if err == nil {
		t.Errorf("expected same length")
	}
}

func TestIncorrectIndex(t *testing.T) {
	m := `1 2
		  	3 4 
		  	5 6`
	matrix, _ := New(m)
	res := matrix.Set(1, 5, 22)
	expected := false

	if res {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}

func TestCorrectIndex(t *testing.T) {
	m := `1 2
		  	3 4 
		  	5 6`
	matrix, _ := New(m)
	res := matrix.Set(1, 1, 22)
	expected := true

	if !res {
		t.Errorf("expected: %t, got %t", expected, res)
	}
}
