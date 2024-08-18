package pkg

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeInteger(t *testing.T) {
	input := "i11e"

	expected := Value{typ: "integer", integer: 11}

	got, _, err := decodeInteger([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, expected.typ, got.typ, "unexpected type")
	assert.Equal(t, expected.integer, got.integer, "unexpected integer value")
}

func TestDecodeString(t *testing.T) {
	input := "15:helicopteraaaaa"

	expected := Value{typ: "string", str: "helicopteraaaaa"}

	got, _, err := decodeString([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, expected.typ, got.typ, "unexpected type")
	assert.Equal(t, expected.str, got.str, "unexpected string value")
}

func TestDecodesList(t *testing.T) {
	input := "li11e10:helicoptere"

	expected := Value{
		typ: "list",
		list: []Value{
			{typ: "integer", integer: 11},
			{typ: "string", str: "helicopter"},
		},
	}

	got, _, err := decodeList([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("decodeList(%q) = %v, want %v", input, got, expected)
	}
}

func TestDecodeDict(t *testing.T) {
	input := "d3:foo3:bar3:bazi42ee"

	expected := Value{
		typ: "dictionary",
		dictionary: map[string]Value{
			"foo": {typ: "string", str: "bar"},
			"baz": {typ: "integer", integer: 42},
		},
	}

	got, _, err := decodeDict([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("decodeDict(%q) = %v, want %v", input, got, expected)
	}
}
