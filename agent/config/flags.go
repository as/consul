package config

import (
	"flag"
	"strconv"
)

// Flags defines the command line flags.
//
// All fields are specified as pointers to simplify merging multiple
// File structures since this allows to determine whether a field has
// been set.
type Flags struct {
	Bootstrap  *bool
	Datacenter *string
}

// File returns the config file representation of the provided
// command line flags.
func (f *Flags) File() File {
	return File{
		Bootstrap:  f.Bootstrap,
		Datacenter: f.Datacenter,
	}
}

func NewFlagSet(f *Flags) *flag.FlagSet {
	fs := flag.NewFlagSet("agent", flag.ContinueOnError)
	fs.Var(newBoolPtrValue(&f.Bootstrap), "bootstrap", "bootstrap yes/no")
	fs.Var(newStringPtrValue(&f.Datacenter), "dc", "datacenter")
	return fs
}

func ParseFlags(args []string) (*Flags, error) {
	f := &Flags{}
	fs := NewFlagSet(f)
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	return f, nil
}

type boolPtrValue struct {
	v **bool
	b bool
}

func newBoolPtrValue(p **bool) *boolPtrValue {
	return &boolPtrValue{p, false}
}

func (s *boolPtrValue) IsBoolFlag() bool { return true }

func (s *boolPtrValue) Set(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*s.v, s.b = &b, true
	return nil
}

func (s *boolPtrValue) Get() interface{} {
	if s.b {
		return *s.v
	}
	return (*bool)(nil)
}

func (s *boolPtrValue) String() string {
	if s.b {
		return strconv.FormatBool(**s.v)
	}
	return ""
}

type stringPtrValue struct {
	v **string
	b bool
}

func newStringPtrValue(p **string) *stringPtrValue {
	return &stringPtrValue{p, false}
}

func (s *stringPtrValue) Set(val string) error {
	*s.v, s.b = &val, true
	return nil
}

func (s *stringPtrValue) Get() interface{} {
	if s.b {
		return *s.v
	}
	return (*string)(nil)
}

func (s *stringPtrValue) String() string {
	if s.b {
		return **s.v
	}
	return ""
}
