package golru

import (
	"fmt"
	"testing"
)

func TestAdd_Get(t *testing.T) {
	c := newBasic(4)
	k := "test"
	v := "fizzbuzz"

	c.Add(k, v)

	stored, ok := c.Get(k)
	if !ok {
		t.Fatalf("expected that adding a key would set the value")
	}

	storedString, _ := stored.(string)
	if storedString != v {
		t.Errorf("unexpected value, expected: %s, got: %s", v, storedString)
	}
}

func TestAdd_Size(t *testing.T) {
	size := 5
	c := newBasic(size)

	for i := 1; i <= size+3; i++ {
		key := fmt.Sprintf("key-%s", i)
		val := fmt.Sprintf("val-%s", i)
		evicted := c.Add(key, val)
		if evicted && i <= size {
			t.Errorf("expected that adding while under the cache would not evict: %d - %d", size, i)
		}

		if !evicted && i > size {
			t.Errorf("expected that adding while at cap would evict: %d - %d", size, i)
		}
	}
}

func TestAdd_ExistingKey(t *testing.T) {
	c := newBasic(1)
	key := "key"
	firstVal := "first"
	secondVal := "second"
	c.Add(key, firstVal)

	evicted := c.Add(key, secondVal)
	if evicted {
		t.Error("expected that adding an existing key would not cause an eviction")
	}

	stored, _ := c.Get(key)
	storedString, _ := stored.(string)

	if storedString != secondVal {
		t.Errorf("got unexpected value from key, expected: %s, got: %s", secondVal, storedString)
	}
}

func TestRemove(t *testing.T) {
	c := newBasic(1)
	key := "key"
	val := "val"
	c.Add(key, val)

	removed := c.Remove(key)
	if !removed {
		t.Fatal("expected that removing a key that exists would return true")
	}

	removed2 := c.Remove(key)
	if removed2 {
		t.Fatal("expected that removing a key that doesn't exist would return false")
	}
}

func TestRecent(t *testing.T) {
	c := newBasic(3)
	c.Add("a", 1)
	c.Add("b", 2)
	c.Add("c", 3)

	// Assert that "a" is removed:
	e := c.Add("d", 4)
	if !e {
		t.Fatal("expected eviction on 4th add")
	}

	_, ok := c.Get("a")
	if ok {
		t.Fatal("expected that the oldest item would be evicted")
	}

	// Access "b", then add, and assert that "b" was not removed,
	// as it was no longer the oldest.
	c.Get("b")
	ev := c.Add("e", 5)
	if !ev {
		t.Fatal("expected eviction on 5th add")
	}

	_, ok = c.Get("b")
	if !ok {
		t.Fatalf("expected that accessing a key would reset it in the order")
	}
}
