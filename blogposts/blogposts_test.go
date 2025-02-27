package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/and-volkov/blogposts"
)

type StubFailingFS struct{}

func (s *StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogposts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromSF(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "Hello\nWorld",
	}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
