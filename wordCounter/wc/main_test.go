package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 3
	res := count(b, true, false)
	if res != exp {
		t.Errorf("expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 a word2\nline 1\n")

	exp := 16
	res := count(b, false, true)
	if res != exp {
		t.Errorf("expecte %d, got %d instead\n", exp, res)
	}
}
