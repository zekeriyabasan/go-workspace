package errorhandling

import "fmt"

type outOfRangeException struct {
	value   int
	message string
}

func (o *outOfRangeException) Error() string {

	return fmt.Sprintf("%d------%s", o.value, o.message)

}
