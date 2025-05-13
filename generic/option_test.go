package generic

import "testing"

func TestOption(t *testing.T) {
	o := NewOption[string]()
	val, err := o.Take()
	if err == nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got: %v", val)
	}

	o.Set("hello friends")
	_, err = o.Take()
	if err != nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got: %v", val)
	}

	o.Clear()
	if o.IsSome() {
		t.Fatal("Option should hava none, but has some")
	}

	{
		defer func() {
			if r := recover(); r != nil {
				t.Log("Expected panic occurred")
			} else {
				t.Error("Expected panic, but none occurred")
			}
		}()
		o.Yank()
	}
}
