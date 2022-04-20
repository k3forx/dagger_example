package greeting

import "testing"

func TestGreeting(t *testing.T) {
	name := "Dagger Test"
	expected := "Hi Dagger Test!"
	value := Greeting(name)

	if expected != value {
		t.Fatalf("Hello(%s) = '%s', expected '%s'", name, value, expected)
	}
}
