package img

import "testing"

func TestJpegconv(t *testing.T) {
	// t.Run("Convert PNG to JPEG", func(t *testing.T) {
	// 	got := Jpegconv(25)
	// 	expected := 50

	// 	assertion(t, got, expected)
	// })
}

func assertion(t testing.TB, got, expected int) {
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}