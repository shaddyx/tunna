package wsTools

import "fmt"

type WsMessage struct {
	EventName *string
	Message   *[]byte
}

const (
	IntSize = 8
)

type Error struct{
	Err string
	error
	NestedError error
}

func (e Error) String() string{
	if e.NestedError != nil{
		return e.String() + ", nested:" + fmt.Sprint("%s", e.NestedError)
	} else {
		return e.String()
	}
}

func (e Error) Error() string {
	return e.Err
}