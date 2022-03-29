package terr

import (
	"fmt"
	"touzi"
)

type ErrInvalidArgument struct {
	Arguments []touzi.Argument
	Position  int
}

func (i *ErrInvalidArgument) Error() string {
	return fmt.Sprintf("invalid argument at position %d: %s", i.Position, string(i.Arguments[i.Position]))
}

func InvalidArgument(argument []touzi.Argument, position int) *ErrInvalidArgument {
	return &ErrInvalidArgument{
		Arguments: argument,
		Position:  position,
	}
}

type ErrInvalidArguments struct {
	Arguments []touzi.Argument
}

func (i *ErrInvalidArguments) Error() string {
	return fmt.Sprintf("invalid arguments: %v", i.Arguments)
}

func InvalidArguments(argument []touzi.Argument) *ErrInvalidArguments {
	return &ErrInvalidArguments{
		Arguments: argument,
	}
}

type ErrTouziNotFound struct {
	Prefix rune
}

func (i *ErrTouziNotFound) Error() string {
	return fmt.Sprintf("touzi with prefix %v not found", i.Prefix)
}

func TouziNotFound(prefix rune) *ErrTouziNotFound {
	return &ErrTouziNotFound{
		Prefix: prefix,
	}
}
