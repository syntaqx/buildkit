package config

// Database defines the database configuration.
type Database struct {
	DSN string
}

// Server defines the webserver configuration.
type Server struct {
	Addr    string
	Port    string
	Debug   bool
	BaseURL string
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string
	Token string
}

// Logging defines logging configuration.
type Logging struct {
	Level  string
	Format string
}

// Tracing defines the tracing client configuration.
type Tracing struct {
	Endpoint string
}

// Config is a combination of all available configurations.
type Config struct {
	Logging  Logging
	Tracing  Tracing
	Server   Server
	Metrics  Metrics
	Database Database
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
