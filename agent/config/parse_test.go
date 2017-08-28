package config

import (
	"strings"
	"testing"

	"github.com/pascaldekloe/goe/verify"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc  string
		fmt   string // json or hcl
		def   File
		files []string
		flags []string
		cfg   Config
		err   error
	}{
		{
			desc: "default config",
			def:  defaultFile,
			cfg:  defaultConfig,
		},

		// cmd line flags
		{
			flags: []string{`-bootstrap`},
			cfg:   Config{Bootstrap: true},
		},
		{
			flags: []string{`-dc`, `a`},
			cfg:   Config{Datacenter: "a"},
		},

		// json cfg file
		{
			fmt:   "json",
			files: []string{`{"bootstrap":true}`},
			cfg:   Config{Bootstrap: true},
		},
		{
			fmt:   "json",
			files: []string{`{"bootstrap":true}`, `{"bootstrap":false}`},
			cfg:   Config{Bootstrap: false},
		},

		// hcl cfg file
		{
			fmt:   "hcl",
			files: []string{`bootstrap=true`},
			cfg:   Config{Bootstrap: true},
		},
		{
			fmt:   "hcl",
			files: []string{`bootstrap=true`, `bootstrap=false`},
			cfg:   Config{Bootstrap: false},
		},

		// precedence rules
		{
			fmt:   "json",
			files: []string{`{"bootstrap":true}`},
			flags: []string{`-bootstrap=false`},
			cfg:   Config{Bootstrap: false},
		},
		{
			fmt:   "hcl",
			files: []string{`bootstrap=true`},
			flags: []string{`-bootstrap=false`},
			cfg:   Config{Bootstrap: false},
		},
	}

	for _, tt := range tests {
		var desc []string
		if tt.desc != "" {
			desc = append(desc, tt.desc)
		}
		if len(tt.files) > 0 {
			s := tt.fmt + ":" + strings.Join(tt.files, ",")
			desc = append(desc, s)
		}
		if len(tt.flags) > 0 {
			s := "flags:" + strings.Join(tt.flags, " ")
			desc = append(desc, s)
		}

		t.Run(strings.Join(desc, ";"), func(t *testing.T) {
			p := &Parser{}

			// start with default config
			p.Files = append(p.Files, tt.def)

			// then add files in order
			for _, s := range tt.files {
				f, err := ParseFile(s)
				if err != nil {
					t.Fatalf("ParseFile failed for %q: %s", s, err)
				}
				p.Files = append(p.Files, f)
			}

			// then add flags
			flags, err := ParseFlags(tt.flags)
			if err != nil {
				t.Fatalf("ParseFlags failed: %s", err)
			}
			p.Files = append(p.Files, flags.File())

			cfg, err := p.Merge()
			if err != nil {
				t.Fatalf("Parse failed: %s", err)
			}

			if !verify.Values(t, "", cfg, tt.cfg) {
				t.FailNow()
			}
		})
	}

}
