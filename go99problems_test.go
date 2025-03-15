package go99problems

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMyLast(t *testing.T) {
	int_input := []int{1, 2, 3, 4}
	int_expected := 4

	int_received, err := MyLast(int_input)
	if err != nil {
		t.Errorf("unexpected error with int slice %v", int_input)
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice %v, expected: %d, received: %d", int_input, int_expected, int_received)
	}

	byte_input := []byte("xyz")
	byte_expected := byte('z')

	byte_received, err := MyLast(byte_input)
	if err != nil {
		t.Errorf("unexpected error with byte slice %v", byte_input)
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice %v, expected: %c, received: %c", byte_input, byte_expected, byte_received)
	}
}

func TestMyButLast(t *testing.T) {
	int_input := []int{1, 2, 3, 4}
	int_expected := 3

	int_received, err := MyButLast(int_input)
	if err != nil {
		t.Errorf("unexpected error with int slice %v", int_input)
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice %v, expected: %d, received: %d", int_input, int_expected, int_received)
	}

	var byte_input []byte

	for i := range 26 {
		byte_input = append(byte_input, byte(i+97))
	}

	byte_expected := byte('y')

	byte_received, err := MyButLast(byte_input)
	if err != nil {
		t.Errorf("unexpected error with byte slice %v", byte_input)
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice %v, expected: %c, received: %c", byte_input, byte_expected, byte_received)
	}
}

func TestElementAt(t *testing.T) {
	int_input := []int{1, 2, 3}
	int_k := 2
	int_expected := 2

	int_received, err := ElementAt(int_input, int_k)
	if err != nil {
		t.Errorf("unexpected error with int slice %v", int_input)
	}

	if int_received != int_expected {
		t.Errorf("fail with int slice %v, expected: %d, received: %d", int_input, int_expected, int_received)
	}

	byte_input := []byte("golang")
	byte_k := 5
	byte_expected := byte('n')

	byte_received, err := ElementAt(byte_input, byte_k)
	if err != nil {
		t.Errorf("unexpected error with byte slice %v", byte_input)
	}

	if byte_received != byte_expected {
		t.Errorf("fail with byte slicei %v, expected: %d, received: %d", byte_input, byte_expected, byte_received)
	}
}

func TestMyLength(t *testing.T) {
	int_input := []int{123, 456, 789}
	int_expected := 3

	int_received := MyLength(int_input)

	if int_received != int_expected {
		t.Errorf("fail with int slice %v, expected: %d, received: %d", int_input, int_expected, int_received)
	}

	byte_input := []byte("Hello, world!")
	byte_expected := 13

	byte_received := MyLength(byte_input)

	if byte_received != byte_expected {
		t.Errorf("fail with byte slice %v, expected: %d, received: %d", byte_input, byte_expected, byte_received)
	}
}

func TestMyReverse(t *testing.T) {
	int_input := []int{1, 2, 3, 4}

	int_expected := make([]int, len(int_input))
	_ = copy(int_expected, int_input)
	slices.Reverse(int_expected)

	int_received := MyReverse(int_input)

	if diff := slices.Compare(int_expected, int_received); diff != 0 {
		t.Errorf("fail with int slice %v, expected: %v, received: %v", int_input, int_expected, int_received)
	}

	byte_input := []byte("A man, a plan, a canal, panama!")

	byte_expected := make([]byte, len(byte_input))
	_ = copy(byte_expected, byte_input)
	slices.Reverse(byte_expected)

	byte_received := MyReverse(byte_input)

	if diff := slices.Compare(byte_expected, byte_received); diff != 0 {
		t.Errorf("fail with byte slice %v, expected: %v, received: %v", byte_input, byte_expected, byte_received)
	}
}

func TestIsPalindrome(t *testing.T) {
	int_false_input := []int{1, 2, 3}
	int_false_received := IsPalindrome(int_false_input)
	int_false_expected := false

	if int_false_received != int_false_expected {
		t.Errorf("fail with int slice %v, expected: %t, received: %t", int_false_input, int_false_received, int_false_expected)
	}

	int_true_input := []int{1, 2, 4, 8, 16, 8, 4, 2, 1}
	int_true_received := IsPalindrome(int_true_input)
	int_true_expected := true

	if int_true_received != int_true_expected {
		t.Errorf("fail with int slice %v, expected: %t, received: %t", int_true_input, int_true_received, int_true_expected)
	}

	byte_true_input := []byte("madamimadam")
	byte_true_received := IsPalindrome(byte_true_input)
	byte_true_expected := true

	if byte_true_received != byte_true_expected {
		t.Errorf("fail with byte slice %v, expected: %t, received: %t", byte_true_input, byte_true_received, byte_true_expected)
	}
}

func TestFlatten(t *testing.T) {
	singleton_input := NestedSlice[int]{
		values:   []int{5},
		children: nil,
	}

	singleton_expected := []int{5}
	singleton_received := Flatten(singleton_input)

	if diff := slices.Compare(singleton_expected, singleton_received); diff != 0 {
		t.Errorf("fail with singleton nested slice %v, expected: %v, received: %v", singleton_input, singleton_expected, singleton_received)
	}

	full_input := NestedSlice[int]{
		values: []int{1, 2},
		children: []*NestedSlice[int]{
			{
				values: []int{3},
				children: []*NestedSlice[int]{
					{values: []int{4, 5}},
				},
			},
		},
	}

	full_expected := []int{1, 2, 3, 4, 5}
	full_received := Flatten(full_input)

	if diff := slices.Compare(full_expected, full_received); diff != 0 {
		t.Errorf("fail with full nested slice %v, expected: %v, received: %v", full_input, full_expected, full_received)
	}

	empty_input := NestedSlice[int]{
		values:   []int{},
		children: nil,
	}

	empty_expected := []int{}
	empty_received := Flatten(empty_input)

	if diff := slices.Compare(empty_expected, empty_received); diff != 0 {
		t.Errorf("fail with empty nested slice %v, expected: %v, received: %v", empty_input, empty_expected, empty_received)
	}
}

func TestCompress(t *testing.T) {
	byte_input := []byte("aaaabccaadeeee")
	byte_expected := []byte("abcade")
	byte_received := Compress(byte_input)

	if diff := slices.Compare(byte_expected, byte_received); diff != 0 {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input, byte_expected, byte_received)
	}

	empty_input := []byte("")
	empty_expected := []byte("")
	empty_received := Compress(empty_input)

	if diff := slices.Compare(empty_expected, empty_received); diff != 0 {
		t.Errorf("fail with empty slice %s, expected: %s, received: %s", empty_input, empty_expected, empty_received)
	}
}

func TestPack(t *testing.T) {
	byte_input := []byte("aaaabccaadeeee")
	byte_expected := [][]byte{[]byte("aaaa"), []byte("b"), []byte("cc"), []byte("aa"), []byte("d"), []byte("eeee")}
	byte_received := Pack(byte_input)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %v, received: %v", byte_input, byte_expected, byte_received)
	}

	empty_input := []byte("")
	empty_expected := []byte("")
	empty_received := Compress(empty_input)

	if diff := slices.Compare(empty_expected, empty_received); diff != 0 {
		t.Errorf("fail with empty slice %s, expected: %s, received: %s", empty_input, empty_expected, empty_received)
	}
}

func TestEncode(t *testing.T) {
	byte_input := []byte("aaaabccaadeeee")
	byte_expected := []RLEPair[byte]{
		{4, byte('a')},
		{1, byte('b')},
		{2, byte('c')},
		{2, byte('a')},
		{1, byte('d')},
		{4, byte('e')},
	}
	byte_received := Encode(byte_input)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %v, received: %v", byte_input, byte_expected, byte_received)
	}
}

func TestEncodeModified(t *testing.T) {
	byte_input := []byte("aaaabccaadeeee")
	byte_expected := [][]any{
		{4, byte('a')},
		{1, byte('b')},
		{2, byte('c')},
		{2, byte('a')},
		{1, byte('d')},
		{4, byte('e')},
	}
	byte_received := EncodeModified(byte_input)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %v, received: %v", byte_input, byte_expected, byte_received)
	}
}

func TestDecode(t *testing.T) {
	byte_input := []RLEPair[byte]{
		{4, byte('a')},
		{1, byte('b')},
		{2, byte('c')},
		{2, byte('a')},
		{1, byte('d')},
		{4, byte('e')},
	}
	byte_expected := []byte("aaaabccaadeeee")
	byte_received := Decode(byte_input)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %v, expected: %s, received: %s", byte_input, byte_expected, byte_received)
	}
}

func TestEncodeDirect(t *testing.T) {
	byte_input := []byte("aaaabccaadeeee")
	byte_expected := []RLEPair[byte]{
		{4, byte('a')},
		{1, byte('b')},
		{2, byte('c')},
		{2, byte('a')},
		{1, byte('d')},
		{4, byte('e')},
	}
	byte_received := EncodeDirect(byte_input)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %v, received: %v", byte_input, byte_expected, byte_received)
	}
}

func TestDupli(t *testing.T) {
	int_input := []int{1, 2, 3}
	int_expected := []int{1, 1, 2, 2, 3, 3}
	int_received := Dupli(int_input)

	if !cmp.Equal(int_expected, int_received) {
		t.Errorf("fail with int slice %v, expected: %v, received: %v", int_input, int_expected, int_received)
	}
}

func TestRepli(t *testing.T) {
	byte_input := []byte("abc")
	byte_expected := []byte("aaabbbccc")
	byte_received := Repli(byte_input, 3)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input, byte_expected, byte_received)
	}
}

func TestDropEvery(t *testing.T) {
	byte_input := []byte("abcdefghik")
	byte_expected := []byte("abdeghk")
	byte_received := DropEvery(byte_input, 3)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input, byte_expected, byte_received)
	}
}

func TestSplit(t *testing.T) {
	byte_input := []byte("abcdefghik")
	byte_expectedleft, byte_expectedright := []byte("abc"), []byte("defghik")
	byte_receivedleft, byte_receivedright := Split(byte_input, 3)

	if !cmp.Equal(byte_expectedleft, byte_receivedleft) || !cmp.Equal(byte_expectedright, byte_receivedright) {
		t.Errorf("fail with byte slice %s, expected left: %s, received left: %s, expected right: %s, expected left: %s", byte_input, byte_expectedleft, byte_receivedleft, byte_expectedright, byte_receivedright)
	}
}

func TestSlice(t *testing.T) {
	byte_input := []byte("abcdefghik")
	byte_expected := []byte("cdefg")
	byte_received := Slice(byte_input, 3, 7)

	if !cmp.Equal(byte_expected, byte_received) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input, byte_expected, byte_received)
	}
}

func TestRotate(t *testing.T) {
	byte_input_pos := []byte("abcdefgh")
	byte_expected_pos := []byte("defghabc")
	byte_received_pos := Rotate(byte_input_pos, 3)

	if !cmp.Equal(byte_expected_pos, byte_received_pos) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input_pos, byte_expected_pos, byte_received_pos)
	}

	byte_input_neg := []byte("abcdefgh")
	byte_expected_neg := []byte("ghabcdef")
	byte_received_neg := Rotate(byte_input_neg, -2)

	if !cmp.Equal(byte_expected_neg, byte_received_neg) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input_neg, byte_expected_neg, byte_received_neg)
	}

	byte_input_zero := []byte("abcdefgh")
	byte_expected_zero := byte_input_zero
	byte_received_zero := Rotate(byte_input_zero, len(byte_input_zero))

	if !cmp.Equal(byte_expected_zero, byte_received_zero) {
		t.Errorf("fail with byte slice %s, expected: %s, received: %s", byte_input_zero, byte_expected_zero, byte_received_zero)
	}
}

func TestRemoveAt(t *testing.T) {
	byte_input := []byte("abcd")
	byte_expected_el, byte_expected_rem := byte('b'), []byte("acd")
	byte_received_el, byte_received_rem := RemoveAt(2, byte_input)

	if !cmp.Equal(byte_expected_el, byte_received_el) || !cmp.Equal(byte_expected_rem, byte_received_rem) {
		t.Errorf("fail with byte slice %s, expected el: %c, received el: %c, expected rem: %s, expected el: %s", byte_input, byte_expected_el, byte_received_el, byte_expected_rem, byte_received_rem)
	}
}
