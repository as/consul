package config

import (
	"time"
)

type Parser struct {
	Files  []File
	Errors []error
}

func (p *Parser) Merge() (Config, error) {
	var c Config

	c.Bootstrap = p.bool(func(f File) *bool { return f.Bootstrap })
	c.CheckUpdateInterval = p.duration(func(f File) *string { return f.CheckUpdateInterval })
	c.Datacenter = p.string(func(f File) *string { return f.Datacenter })

	return c, nil
}

func (p *Parser) bool(val func(File) *bool) bool {
	var v bool
	for _, f := range p.Files {
		if x := val(f); x != nil {
			v = *x
		}
	}
	return v
}

func (p *Parser) duration(val func(File) *string) time.Duration {
	var v string
	for _, f := range p.Files {
		if x := val(f); x != nil {
			v = *x
		}
	}

	// todo(fs): parse float and time.Duration really???
	d, err := time.ParseDuration(v)
	if err != nil {
		p.Errors = append(p.Errors, err)
		return 0
	}
	return d
}

func (p *Parser) string(val func(File) *string) string {
	var v string
	for _, f := range p.Files {
		if x := val(f); x != nil {
			v = *x
		}
	}
	return v
}
