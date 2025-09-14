package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")

	expected := 4
	actual := count(b, Flags{isForLines: false, isForBytes: false})

	if actual != expected {
		t.Errorf("Expected [%d] words, got %d instead.\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("line1\nline2\nlin3")

	expected := 3
	actual := count(b, Flags{isForLines: true, isForBytes: false})

	if actual != expected {
		t.Errorf("Expected [%d] lines, got %d instead.\n", expected, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("fox")

	expected := 3
	actual := count(b, Flags{isForLines: false, isForBytes: true})

	if actual != expected {
		t.Errorf("Expected [%d] bytes, got %d instead.\n", expected, actual)
	}
}
