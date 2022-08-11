package concurrency_test

import (
	concurrency "10-concurrency"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
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

func BenchmarkCheckWebsites(b *testing.B) {
	// ?? SLOW-ASS website checker 😈
	slowWc := func(_ string) bool {
		time.Sleep(20 * time.Millisecond)
		return true
	}
	urls := generateUrls(100)

	// ?? reset time-spend from previous statements before start measuring
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowWc, urls)
	}
}

func generateUrls(total int) (urls []string) {
	urls = make([]string, total)
	digits := fmt.Sprintf("%d", len(strconv.Itoa(total)))

	for i := 0; i < len(urls); i++ {
		n := fmt.Sprintf("%0"+digits+"d", i+1)
		urls[i] = "GENERATED_URL_" + n
	}

	return
}
