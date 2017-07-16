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
