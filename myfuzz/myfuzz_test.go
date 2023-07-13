package myfuzz

import (
	"bytes"
	"encoding/hex"
	"math"
	"net/url"
	"reflect"
	"testing"
)

func FuzzFoo(f *testing.F) {
	f.Add(5, "hello")
	f.Fuzz(func(t *testing.T, i int, s string) {
		out, err := Foo(i, s)
		if err != nil && out != "" {
			t.Errorf("%q, %v", out, err)
		}
	})
}

func FuzzDontSayBye(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, s string) {
		err := DontSayBye(s)
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}

func FuzzDontSayGoodbye(f *testing.F) {
	f.Add("goodbye")
	f.Fuzz(func(t *testing.T, s string) {
		err := DontSayGoodbye(s)
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}

func FuzzParseQuery(f *testing.F) {
	f.Add("x=1&y=2")
	f.Fuzz(func(t *testing.T, queryStr string) {
		query, err := url.ParseQuery(queryStr)
		if err != nil {
			t.Skip()
		}
		queryStr2 := query.Encode()
		query2, err := url.ParseQuery(queryStr2)
		if err != nil {
			t.Fatalf("ParseQuery failed to decode a valid encoded query %s: %v", queryStr2, err)
		}
		if !reflect.DeepEqual(query, query2) {
			t.Errorf("ParseQuery gave different query after being encoded\nbefore: %v\nafter: %v", query, query2)
		}
	})
}

func FuzzAbs(f *testing.F) {
	f.Add(99)
	f.Fuzz(func(t *testing.T, i int) {
		if Abs(i) != int(math.Abs(float64(i))) {
			t.Errorf("%v", i)
		}
	})
}

func FuzzHex(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}
