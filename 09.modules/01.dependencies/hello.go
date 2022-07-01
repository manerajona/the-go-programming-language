package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Quote returns a quote in latest version
func Quote() string {
	return quote.Glass()
}

// Proverb returns a quote version 3
func Proverb() string {
	return quoteV3.Concurrency()
}
