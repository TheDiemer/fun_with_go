package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eric", "")
		want := "Hello, Eric"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Nicolaus", "French")
		want := "Bonjour, Nicolaus"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Karl", "German")
		want := "Hallo, Karl"
		assertCorrectMessage(t, got, want)
	})
}
