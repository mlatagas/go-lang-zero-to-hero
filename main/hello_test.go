package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people in English", func(t *testing.T) {

		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello to people in Zulu", func(t *testing.T) {

		got := Hello("Chris", "zulu")
		want := "Sawubona, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello to someone in Spanish", func(t *testing.T) {

		got := Hello("Chris", "spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
