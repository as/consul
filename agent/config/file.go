package config

import "github.com/hashicorp/hcl"

// File defines the format of a config file.
//
// All fields are specified as pointers to simplify merging multiple
// File structures since this allows to determine whether a field has
// been set.
type File struct {
	Bootstrap           *bool
	CheckUpdateInterval *string
	Datacenter          *string
}

func ParseFile(s string) (File, error) {
	var f File
	if err := hcl.Decode(&f, s); err != nil {
		return File{}, err
	}
	return f, nil
}
