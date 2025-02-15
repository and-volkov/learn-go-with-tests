package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furfurfur.furfur"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("check websites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furfurfur.furfur",
		}

		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furfurfur.furfur":    false,
		}
		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}
