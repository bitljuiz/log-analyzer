package flags

import (
	"errors"
)

// Flag это обертка над флагом консольной утилиты.
type Flag struct {
	Name          string
	ShorthandName string
	Use           string
	Value         Value
}

var (
	ErrCannotGetValue = errors.New("cannot get flag value")
)

func (f *Flag) GetString() (string, error) {
	switch val := f.Value.(type) {
	case *StringValue:
		return val.Value(), nil
	default:
		return "", ErrCannotGetValue
	}
}

func (f *Flag) GetInt() (int, error) {
	switch val := f.Value.(type) {
	case *IntegerValue:
		return val.Value(), nil
	default:
		return -1, ErrCannotGetValue
	}
}

type Value interface {
	Type() string
}

type StringValue struct {
	value        string
	defaultValue string
}

func NewStringValue(defaultValue string) *StringValue {
	s := StringValue{
		defaultValue: defaultValue,
	}

	return &s
}

func (sv *StringValue) Type() string { return "string" }

func (sv *StringValue) Pointer() *string {
	return &sv.value
}

func (sv *StringValue) Value() string {
	return sv.value
}

func (sv *StringValue) DefaultValue() string {
	return sv.defaultValue
}

type IntegerValue struct {
	value        int
	defaultValue int
}

func NewIntegerValue(defaultValue int) *IntegerValue {
	s := IntegerValue{
		defaultValue: defaultValue,
	}

	return &s
}

func (iv *IntegerValue) Type() string { return "int" }

func (iv *IntegerValue) Pointer() *int {
	return &iv.value
}

func (iv *IntegerValue) Value() int {
	return iv.value
}

func (iv *IntegerValue) DefaultValue() int {
	return iv.defaultValue
}
