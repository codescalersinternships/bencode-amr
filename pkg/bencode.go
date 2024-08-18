package pkg

import (
	"fmt"
	"strconv"
)

type Value struct {
	typ        string
	integer    int
	str        string
	list       []Value
	dictionary map[string]Value
}

func Decode(s []byte) (Value, []byte, error) {
	switch _type := s[0]; {
	case _type == 'i':
		return decodeInteger(s)
	case _type >= '0' && _type <= '9':
		return decodeString(s)
	case _type == 'd':
		return decodeDict(s)
	case _type == 'l':
		return decodeList(s)
	default:
		return Value{}, nil, fmt.Errorf("invalid starting character: %s", s)
	}
}

func decodeInteger(s []byte) (Value, []byte, error) {
	end := 0
	for i := 1; i < len(s); i++ {
		if s[i] == 'e' {
			end = i
			break
		}
	}
	if end == 0 {
		return Value{}, nil, fmt.Errorf("invalid end character: %v", s)
	}

	v := Value{typ: "integer"}
	var err error
	v.integer, err = strconv.Atoi(string(s[1:end]))
	if err != nil {
		return Value{}, nil, err
	}

	return v, s[end+1:], nil
}

func decodeString(s []byte) (Value, []byte, error) {
	idx, err := getFirstByte(s, ':')
	if err != nil {
		return Value{}, nil, err
	}

	l, err := strconv.Atoi(string(s[:idx]))
	if err != nil {
		return Value{}, nil, err
	}

	v := Value{typ: "string", str: string(s[idx+1 : idx+1+l])}
	return v, s[idx+1+l:], nil
}

func decodeList(s []byte) (Value, []byte, error) {
	v := Value{typ: "list"}
	v.list = make([]Value, 0)

	s = s[1:]

	for len(s) > 0 {
		if s[0] == 'e' {
			s = s[1:]
			break
		}

		val, remainder, err := Decode(s)
		if err != nil {
			return Value{}, nil, err
		}
		v.list = append(v.list, val)

		s = remainder
	}

	return v, s, nil
}

func decodeDict(s []byte) (Value, []byte, error) {
	v := Value{typ: "dictionary"}
	v.dictionary = make(map[string]Value)

	s = s[1:]

	for len(s) > 0 {
		if s[0] == 'e' {
			s = s[1:]
			break
		}

		key, remainder, err := decodeString(s)
		if err != nil {
			return Value{}, nil, err
		}

		value, remainder, err := Decode(remainder)
		if err != nil {
			return Value{}, nil, err
		}

		v.dictionary[key.str] = value

		s = remainder
	}

	return v, s, nil
}

func getFirstByte(s []byte, b byte) (int, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i, nil
		}
	}
	return 0, fmt.Errorf("invalid input: %s", s)
}
