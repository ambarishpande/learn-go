package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	output := Hello()
	if output != "Hello" {
		t.Errorf("Output mismatch: Expected Hello but got %v", output)
	}
}

func TestQuote(t *testing.T) {
	output := Quote()
	if output != "Hello, world." {
		t.Errorf("Output mismatch: Expected Hello but got %v", output)
	}
}
