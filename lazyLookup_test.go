package golru

import (
	"fmt"
	"testing"
)

func makeVal(s string) string {
	return fmt.Sprintf("val-%s", s)
}

func TestLazyLookup(t *testing.T) {
	var lookupCalled bool

	lookup := func(k interface{}) (interface{}, bool) {
		lookupCalled = true
		s := k.(string)
		return makeVal(s), true
	}

	c := newLazyLookup(3, lookup)
	k := "key"
	v := makeVal(k)

	stored, ok := c.Get(k)
	if !ok {
		t.Fatal("expected that lazy cache could find value for a get miss")
	}

	if !lookupCalled {
		t.Fatal("expected that lazy cache would call lookup func on get miss")
	}

	if stored != v {
		t.Fatalf("got unexpected cache value, expected: %s, got: %s", v, stored)
	}
}
