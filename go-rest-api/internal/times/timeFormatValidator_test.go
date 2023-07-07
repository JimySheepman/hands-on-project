package times

import "testing"

func TestTimeFormatValidator(t *testing.T) {
	t.Run("regex is succeed", func(t *testing.T) {
		got := TimeFormatValidator("2016-08-02")
		want := false
		if got != want {
			t.Errorf("Time validation is not succeed: got %v want %v", got, want)
		}

	})

	t.Run("regex is not succeed", func(t *testing.T) {
		got := TimeFormatValidator("2016-8-02")
		want := true
		if got != want {
			t.Errorf("Time validation is not succeed: got %v want %v", got, want)
		}

	})
}
