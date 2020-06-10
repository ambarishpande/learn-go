package hello

import "rsc.io/quote"

// Hello : Method to Say Hello
func Hello() string {
	return "Hello"
}

func Quote() string {
	return quote.Hello()
}
