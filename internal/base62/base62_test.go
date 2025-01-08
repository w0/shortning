package base62

import "testing"

func TestEncode(t *testing.T) {
	m := map[int]string{
		123:    "1z",
		413:    "6f",
		4:      "4",
		234879: "z6N",
		31:     "V",
	}

	for k, v := range m {
		res := Encode(k)

		if res != v {
			t.Fatalf("got %s, want %s", res, v)
		}
	}

}

func TestDecode(t *testing.T) {
	m := map[string]int{
		"/1z":  123,
		"/6f":  413,
		"/4":   4,
		"/z6N": 234879,
		"/V":   31,
	}

	for k, v := range m {
		res := Decode(k)

		if res != v {
			t.Fatalf("got %d, want %d", res, v)
		}
	}
}
