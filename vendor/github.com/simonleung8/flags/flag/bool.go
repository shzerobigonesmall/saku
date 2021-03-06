package cliFlags

import "strconv"

type BoolFlag struct {
	Name      string
	Value     bool
	Usage     string
	ShortName string
}

func (f *BoolFlag) Set(v string) {
	b, _ := strconv.ParseBool(v)
	f.Value = b
}

func (f *BoolFlag) String() string {
	return f.Usage
}

func (f *BoolFlag) GetName() string {
	return f.Name
}

func (f *BoolFlag) GetShortName() string {
	return f.ShortName
}

func (f *BoolFlag) GetValue() interface{} {
	return f.Value
}
