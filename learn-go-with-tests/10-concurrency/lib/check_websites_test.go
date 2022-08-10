package concurrency_test

import (
	concurrency "10-concurrency/lib"
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {
	wc := func(url string) bool {
		if url == "waat://furhurterwe.geds" {
			return false
		}
		return true
	}
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	got := concurrency.CheckWebsites(wc, websites)
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
