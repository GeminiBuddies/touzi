package touzi

import (
	"fmt"
)

type ErrInvalidArgument struct {
	Arguments []Argument
	Position  int
}

func (i *ErrInvalidArgument) Error() string {
	return fmt.Sprintf("invalid argument at position %d: %s", i.Position, string(i.Arguments[i.Position]))
}

func ErrorInvalidArgument(argument []Argument, position int) *ErrInvalidArgument {
	return &ErrInvalidArgument{
		Arguments: argument,
		Position:  position,
	}
}

type ErrInvalidArguments struct {
	Arguments []Argument
}

func (i *ErrInvalidArguments) Error() string {
	return fmt.Sprintf("invalid arguments: %v", i.Arguments)
}

func ErrorInvalidArguments(argument []Argument) *ErrInvalidArguments {
	return &ErrInvalidArguments{
		Arguments: argument,
	}
}

type ErrTouziNotFound struct {
	Prefix rune
}

func (i *ErrTouziNotFound) Error() string {
	return fmt.Sprintf("touzi with prefix '%c' not found", i.Prefix)
}

func ErrorTouziNotFound(prefix rune) *ErrTouziNotFound {
	return &ErrTouziNotFound{
		Prefix: prefix,
	}
}
