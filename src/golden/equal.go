package golden

import (
	"encoding/json"
	"fmt"
	"io"
)

// Equal returns whether JSON data parsed from two readers are equal. Stops
// reading once equality is known. Returns an error if either of the readers
// returned an unexpected error.
func Equal(x, y io.Reader) (equal bool, err error) {
	dx := json.NewDecoder(x)
	dy := json.NewDecoder(y)
	for {
		tx, ex := dx.Token()
		ty, ey := dy.Token()
		if ex != ey {
			switch {
			case ex != nil:
				return false, fmt.Errorf("x error: %w", ex)
			case ey != nil:
				return false, fmt.Errorf("y error: %w", ey)
			}
		}
		if ex == io.EOF {
			break
		}
		if ex != nil {
			return false, ex
		}
		if tx != ty {
			return false, nil
		}
	}
	return true, nil
}
