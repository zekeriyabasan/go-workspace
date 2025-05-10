package errorhandling

import "fmt"

type invalidValue struct {
	message string
}

func (v *invalidValue) Error() string {

	return fmt.Sprint(v.message)

}
