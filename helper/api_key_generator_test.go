package helper

import (
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	if key := GenerateApiKey(32); key == "" {
		t.Error("Key is empty")
	} else {
		if len(key)%4 != 0 {
			t.Error("Key is not a multiple of 4")
			t.Log(len(key))
		}
	}

}
