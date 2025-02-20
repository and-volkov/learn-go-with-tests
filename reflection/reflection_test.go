package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Andrey", "Moscow"},
			[]string{"Andrey", "Moscow"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Andrey", 34},
			[]string{"Andrey"},
		},
		{
			"nested struct",
			Person{
				"Andrey",
				Profile{34, "Moscow"},
			},
			[]string{"Andrey", "Moscow"},
		},
		{
			"pointers to things",
			&Person{
				"Andrey",
				Profile{34, "Moscow"},
			},
			[]string{"Andrey", "Moscow"},
		},
		{
			"slices",
			[]Profile{
				{33, "Moscow"},
				{34, "Tbilisi"},
			},
			[]string{"Moscow", "Tbilisi"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Moscow"},
				{34, "Tbilisi"},
			},
			[]string{"Moscow", "Tbilisi"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow": "Moo",
			"Dog": "Woof",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Woof")
	})
	t.Run("with channels", func(t *testing.T) {
		aChan := make(chan Profile)

		go func() {
			aChan <- Profile{33, "Moscow"}
			aChan <- Profile{34, "Tbilisi"}
			close(aChan)
		}()

		var got []string
		want := []string{"Moscow", "Tbilisi"}

		walk(aChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with func", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Moscow"}, Profile{34, "Tbilisi"}
		}

		var got []string
		want := []string{"Moscow", "Tbilisi"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v contain %q, but it  didn't", haystack, needle)
	}
}
