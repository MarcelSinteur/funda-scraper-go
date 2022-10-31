package config

import (
	"github.com/CloudyKit/jet"
)

// AppConfig holds the application config
type AppConfig struct {
	BaseUrl      string
	InProduction bool
	View         *jet.Set
}
