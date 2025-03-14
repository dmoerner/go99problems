package go99problems

import "testing"

func TestMyLast(t *testing.T) {
	int_input := []int{1, 2, 3, 4}
	int_expected := 4

	int_received := MyLast(int_input)

	if int_received != int_expected {
		t.Errorf("fail with int slice, expected: %d, received: %d", int_expected, int_received)
	}

	byte_input := []byte("xyz")
	byte_expected := byte('z')

	byte_received := MyLast(byte_input)

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice, expected: %b, received: %b", byte_expected, byte_received)
	}
}
