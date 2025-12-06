package pokecache

import (
	"testing"
	"time"
)

type answer struct {
	val     []byte
	isFound bool
}

func TestAddGet(t *testing.T) {
	type test struct {
		name  string
		input struct {
			key string
			val []byte
		}
		want answer
	}

	tests := []test{
		{name: "simple",
			input: struct {
				key string
				val []byte
			}{
				key: "www.example.com",
				val: []byte("imjusttesting"),
			}, want: answer{
				val:     []byte("imjusttesting"),
				isFound: true,
			},
		},
	}

	for _, tc := range tests {
		cache := NewCache(5 * time.Second)
		cache.Add(tc.input.key, tc.input.val)
		val, ok := cache.Get(tc.input.key)
		if !ok {
			t.Fatalf("expected to find key")
		}
		if string(val) != string(tc.want.val) {
			t.Fatalf("expected to find value")
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
