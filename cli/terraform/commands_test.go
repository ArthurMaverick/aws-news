package terraform

import "testing"

func TestIfParametersIsEmpty(t *testing.T) {

	cmd := NewTFCmd("test", "test", "test")

	if cmd.Backend.Bucket == "" {
		t.Error("Bucket is empty")
	}
	if cmd.Backend.Key == "" {
		t.Error("Key is empty")
	}
	if cmd.Backend.Region == "" {
		t.Error("Region is empty")
	}
}
