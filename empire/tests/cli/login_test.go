package cli_test

import (
	"strings"
	"testing"

	"github.com/remind101/empire/empire/empiretest"
)

func TestLogin(t *testing.T) {
	e := empiretest.NewEmpire(t)
	s := empiretest.NewServer(t, e)
	defer s.Close()

	input := "fake\nbar\n"

	cmd := NewCmd(s.URL, "login")
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(out), "Enter email: Logged in.\n"; got != want {
		t.Fatalf("%q", got)
	}
}

func TestLoginUnauthorized(t *testing.T) {
	e := empiretest.NewEmpire(t)
	s := empiretest.NewServer(t, e)
	defer s.Close()

	input := "foo\nbar\n"

	cmd := NewCmd(s.URL, "login")
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected an error")
	}

	if got, want := string(out), "Enter email: error: Request not authenticated, API token is missing, invalid or expired Log in with `emp login`.\n"; got != want {
		t.Fatalf("%q", got)
	}
}

func TestLoginTwoFactor(t *testing.T) {
	e := empiretest.NewEmpire(t)
	s := empiretest.NewServer(t, e)
	defer s.Close()

	input := "twofactor\nbar\ncode\n"

	cmd := NewCmd(s.URL, "login")
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(out), "Enter email: Enter two-factor auth code: Logged in.\n"; got != want {
		t.Fatalf("%q", got)
	}
}
