package go99problems

import (
	"slices"
	"testing"
)

func TestMyLast(t *testing.T) {
	int_input := []int{1, 2, 3, 4}
	int_expected := 4

	int_received, err := MyLast(int_input)
	if err != nil {
		t.Errorf("unexpected error with int slice")
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice, expected: %d, received: %d", int_expected, int_received)
	}

	byte_input := []byte("xyz")
	byte_expected := byte('z')

	byte_received, err := MyLast(byte_input)
	if err != nil {
		t.Errorf("unexpected error with byte slice")
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice, expected: %c, received: %c", byte_expected, byte_received)
	}
}

func TestMyButLast(t *testing.T) {
	int_input := []int{1, 2, 3, 4}
	int_expected := 3

	int_received, err := MyButLast(int_input)
	if err != nil {
		t.Errorf("unexpected error with int slice")
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice, expected: %d, received: %d", int_expected, int_received)
	}

	var byte_input []byte

	for i := range 26 {
		byte_input = append(byte_input, byte(i+97))
	}

	byte_expected := byte('y')

	byte_received, err := MyButLast(byte_input)
	if err != nil {
		t.Errorf("unexpected error with byte slice")
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice, expected: %c, received: %c", byte_expected, byte_received)
	}
}

func TestElementAt(t *testing.T) {
	int_input := []int{1, 2, 3}
	int_k := 2
	int_expected := 2

	int_received, err := ElementAt(int_input, int_k)
	if err != nil {
		t.Errorf("unexpected error with int slice")
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice, expected: %d, received: %d", int_expected, int_received)
	}

	byte_input := []byte("golang")
	byte_k := 5
	byte_expected := byte('n')

	byte_received, err := ElementAt(byte_input, byte_k)
	if err != nil {
		t.Errorf("unexpected error with byte slice")
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice, expected: %d, received: %d", byte_expected, byte_received)
	}
}

func TestMyLength(t *testing.T) {
	int_input := []int{123, 456, 789}
	int_expected := 3

	int_received := MyLength(int_input)

	if int_received != int_expected {
		t.Errorf("fail with int slice, expected: %d, received: %d", int_expected, int_received)
	}

	byte_input := []byte("Hello, world!")
	byte_expected := 13

	byte_received := MyLength(byte_input)

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice, expected: %d, received: %d", byte_expected, byte_received)
	}
}

func TestMyReverse(t *testing.T) {
	int_input := []int{1, 2, 3, 4}

	int_expected := make([]int, len(int_input))
	_ = copy(int_expected, int_input)
	slices.Reverse(int_expected)

	int_received := MyReverse(int_input)

	if diff := slices.Compare(int_expected, int_received); diff != 0 {
		t.Errorf("fail with int slice, expected: %v, received: %v", int_expected, int_received)
	}

	byte_input := []byte("A man, a plan, a canal, panama!")

	byte_expected := make([]byte, len(byte_input))
	_ = copy(byte_expected, byte_input)
	slices.Reverse(byte_expected)

	byte_received := MyReverse(byte_input)

	if diff := slices.Compare(byte_expected, byte_received); diff != 0 {
		t.Errorf("fail with byte slice, expected: %v, received: %v", byte_expected, byte_received)
	}
}
