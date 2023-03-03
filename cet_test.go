package gocet_test

import (
	"testing"
)

func TestCalculateCET(t *testing.T) {

	t.Run("It should not return an error", func(t *testing.T) {

		cet := CalculateCET(400, 100, 200, 12)
		expected := 948.98

		if expected != cet {
			t.Errorf("expected %f. Got %f", expected, cet)
		}
	})

}
