package moment

import "testing"

func TestCatchThrowReturnsThrownError(t *testing.T) {
	f := func() (err error) {
		defer catchThrow(&err)
		panic(typeError("Cannot read properties of null (reading '0')"))
	}
	te, ok := f().(*ThrowError)
	if !ok || te.Name != "TypeError" || te.Message != "Cannot read properties of null (reading '0')" {
		t.Fatalf("got %#v", te)
	}
}

func TestCatchThrowRepanicsOtherValues(t *testing.T) {
	defer func() {
		if r := recover(); r != "bug" {
			t.Fatalf("recovered %#v, want the original panic", r)
		}
	}()
	func() (err error) {
		defer catchThrow(&err)
		panic("bug")
	}()
}

func TestConstructorsValid(t *testing.T) {
	if m, err := New("2024-01-02"); err != nil || !m.IsValid() {
		t.Fatalf("New: %v", err)
	}
	if m, err := CreateUTC("2024-01-02"); err != nil || !m.IsValid() {
		t.Fatalf("CreateUTC: %v", err)
	}
	if _, err := CallStatic("moment", []Value{"2024-01-02"}); err != nil {
		t.Fatalf("CallStatic: %v", err)
	}
}
