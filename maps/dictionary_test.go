package mymaps

import "testing"

// func TestSearch(t *testing.T) {
// 	dictionary := map[string]string{"test": "this is just a test"}

// 	got := Search(dictionary, "test")
// 	want := "this is just a test"

// 	if got != want {
// 		t.Errorf("Got %q want %q", got, want)
// 	}
// }

// refactor to below

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"
	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
