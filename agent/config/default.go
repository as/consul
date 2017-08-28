package config

import "time"

func pBool(v bool) *bool                       { return &v }
func pString(v string) *string                 { return &v }
func pDuration(v time.Duration) *time.Duration { return &v }

// defaultFile is the default configuration file.
var defaultFile = File{
	Bootstrap:           pBool(false),
	CheckUpdateInterval: pString("5m"),
	Datacenter:          pString("dc1"),
}

// defaultConfig is the default runtime configuration which must
// be identical from merging the defaultFile into a configuration.
var defaultConfig = Config{
	Bootstrap:           false,
	CheckUpdateInterval: 5 * time.Minute,
	Datacenter:          "dc1",
}
