package config

import (
	"time"
)

// Config is the runtime configuration.
type Config struct {
	Bootstrap           bool
	CheckUpdateInterval time.Duration
	Datacenter          string
}
