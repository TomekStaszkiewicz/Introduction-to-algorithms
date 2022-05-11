package dataTypes

import (
	"testing"
	"bytes"
)

func TestAlwaysSameRootResult(t *testing.T) {
	data := [][]byte{ []byte("Hello world"), []byte("Mamma mia"), []byte("golang"), []byte("someothertextvalue")}
	val := GetRoot(data)
	
	if bytes.Compare(val, GetRoot(data)) != 0 {
		t.Fatalf("Wrong result value")
	}
	
}
