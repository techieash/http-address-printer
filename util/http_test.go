package util

import (
	"fmt"
	"log"
	"testing"
)

func TestBoundedParallelRequest(t *testing.T) {
	urls := []string{"http://google.com", "http://adjust.com"}
	results := BoundedParallelGet(urls, 3)
	fmt.Println(results)
}

func TestBoundedParallelRequestWithInvalidURL(t *testing.T) {
	urls := []string{"google.com", "http://adjust.com"}
	results := BoundedParallelGet(urls, 3)
	log.Fatal(results)
}
