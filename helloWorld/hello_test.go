package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Human", "")
		want := "What up, Human!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'What up!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "What up!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Human", "Portuguese")
		want := "Olá, Human!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Human", "French")
		want := "Bonjour, Human!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// when a test fails,
	// the line number reported will be in our function call rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
